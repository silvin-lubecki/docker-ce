package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/docker/distribution/uuid"
	"github.com/docker/docker/api"
	apiserver "github.com/docker/docker/api/server"
	buildbackend "github.com/docker/docker/api/server/backend/build"
	"github.com/docker/docker/api/server/middleware"
	"github.com/docker/docker/api/server/router"
	"github.com/docker/docker/api/server/router/build"
	checkpointrouter "github.com/docker/docker/api/server/router/checkpoint"
	"github.com/docker/docker/api/server/router/container"
	distributionrouter "github.com/docker/docker/api/server/router/distribution"
	"github.com/docker/docker/api/server/router/image"
	"github.com/docker/docker/api/server/router/network"
	pluginrouter "github.com/docker/docker/api/server/router/plugin"
	swarmrouter "github.com/docker/docker/api/server/router/swarm"
	systemrouter "github.com/docker/docker/api/server/router/system"
	"github.com/docker/docker/api/server/router/volume"
	"github.com/docker/docker/cli/debug"
	cliflags "github.com/docker/docker/cli/flags"
	"github.com/docker/docker/daemon"
	"github.com/docker/docker/daemon/cluster"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/daemon/logger"
	"github.com/docker/docker/dockerversion"
	"github.com/docker/docker/libcontainerd"
	dopts "github.com/docker/docker/opts"
	"github.com/docker/docker/pkg/authorization"
	"github.com/docker/docker/pkg/jsonlog"
	"github.com/docker/docker/pkg/listeners"
	"github.com/docker/docker/pkg/pidfile"
	"github.com/docker/docker/pkg/plugingetter"
	"github.com/docker/docker/pkg/signal"
	"github.com/docker/docker/plugin"
	"github.com/docker/docker/registry"
	"github.com/docker/docker/runconfig"
	"github.com/docker/go-connections/tlsconfig"
	swarmapi "github.com/docker/swarmkit/api"
	"github.com/spf13/pflag"
)

// DaemonCli represents the daemon CLI.
type DaemonCli struct {
	*config.Config
	configFile *string
	flags      *pflag.FlagSet

	api             *apiserver.Server
	d               *daemon.Daemon
	authzMiddleware *authorization.Middleware // authzMiddleware enables to dynamically reload the authorization plugins
}

// NewDaemonCli returns a daemon CLI
func NewDaemonCli() *DaemonCli {
	return &DaemonCli{}
}

func (cli *DaemonCli) start(opts daemonOptions) (err error) {
	stopc := make(chan bool)
	defer close(stopc)

	// warn from uuid package when running the daemon
	uuid.Loggerf = logrus.Warnf

	opts.common.SetDefaultOptions(opts.flags)

	if cli.Config, err = loadDaemonCliConfig(opts); err != nil {
		return err
	}
	cli.configFile = &opts.configFile
	cli.flags = opts.flags

	if cli.Config.Debug {
		debug.Enable()
	}

	if cli.Config.Experimental {
		logrus.Warn("Running experimental build")
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: jsonlog.RFC3339NanoFixed,
		DisableColors:   cli.Config.RawLogs,
	})

	if err := setDefaultUmask(); err != nil {
		return fmt.Errorf("Failed to set umask: %v", err)
	}

	if len(cli.LogConfig.Config) > 0 {
		if err := logger.ValidateLogOpts(cli.LogConfig.Type, cli.LogConfig.Config); err != nil {
			return fmt.Errorf("Failed to set log opts: %v", err)
		}
	}

	// Create the daemon root before we create ANY other files (PID, or migrate keys)
	// to ensure the appropriate ACL is set (particularly relevant on Windows)
	if err := daemon.CreateDaemonRoot(cli.Config); err != nil {
		return err
	}

	if cli.Pidfile != "" {
		pf, err := pidfile.New(cli.Pidfile)
		if err != nil {
			return fmt.Errorf("Error starting daemon: %v", err)
		}
		defer func() {
			if err := pf.Remove(); err != nil {
				logrus.Error(err)
			}
		}()
	}

	serverConfig := &apiserver.Config{
		Logging:     true,
		SocketGroup: cli.Config.SocketGroup,
		Version:     dockerversion.Version,
		EnableCors:  cli.Config.EnableCors,
		CorsHeaders: cli.Config.CorsHeaders,
	}

	if cli.Config.TLS {
		tlsOptions := tlsconfig.Options{
			CAFile:             cli.Config.CommonTLSOptions.CAFile,
			CertFile:           cli.Config.CommonTLSOptions.CertFile,
			KeyFile:            cli.Config.CommonTLSOptions.KeyFile,
			ExclusiveRootPools: true,
		}

		if cli.Config.TLSVerify {
			// server requires and verifies client's certificate
			tlsOptions.ClientAuth = tls.RequireAndVerifyClientCert
		}
		tlsConfig, err := tlsconfig.Server(tlsOptions)
		if err != nil {
			return err
		}
		serverConfig.TLSConfig = tlsConfig
	}

	if len(cli.Config.Hosts) == 0 {
		cli.Config.Hosts = make([]string, 1)
	}

	api := apiserver.New(serverConfig)
	cli.api = api

	var hosts []string

	for i := 0; i < len(cli.Config.Hosts); i++ {
		var err error
		if cli.Config.Hosts[i], err = dopts.ParseHost(cli.Config.TLS, cli.Config.Hosts[i]); err != nil {
			return fmt.Errorf("error parsing -H %s : %v", cli.Config.Hosts[i], err)
		}

		protoAddr := cli.Config.Hosts[i]
		protoAddrParts := strings.SplitN(protoAddr, "://", 2)
		if len(protoAddrParts) != 2 {
			return fmt.Errorf("bad format %s, expected PROTO://ADDR", protoAddr)
		}

		proto := protoAddrParts[0]
		addr := protoAddrParts[1]

		// It's a bad idea to bind to TCP without tlsverify.
		if proto == "tcp" && (serverConfig.TLSConfig == nil || serverConfig.TLSConfig.ClientAuth != tls.RequireAndVerifyClientCert) {
			logrus.Warn("[!] DON'T BIND ON ANY IP ADDRESS WITHOUT setting --tlsverify IF YOU DON'T KNOW WHAT YOU'RE DOING [!]")
		}
		ls, err := listeners.Init(proto, addr, serverConfig.SocketGroup, serverConfig.TLSConfig)
		if err != nil {
			return err
		}
		ls = wrapListeners(proto, ls)
		// If we're binding to a TCP port, make sure that a container doesn't try to use it.
		if proto == "tcp" {
			if err := allocateDaemonPort(addr); err != nil {
				return err
			}
		}
		logrus.Debugf("Listener created for HTTP on %s (%s)", proto, addr)
		hosts = append(hosts, protoAddrParts[1])
		api.Accept(addr, ls...)
	}

	registryService := registry.NewService(cli.Config.ServiceOptions)
	containerdRemote, err := libcontainerd.New(cli.getLibcontainerdRoot(), cli.getPlatformRemoteOptions()...)
	if err != nil {
		return err
	}
	signal.Trap(func() {
		cli.stop()
		<-stopc // wait for daemonCli.start() to return
	})

	// Notify that the API is active, but before daemon is set up.
	preNotifySystem()

	pluginStore := plugin.NewStore()

	if err := cli.initMiddlewares(api, serverConfig, pluginStore); err != nil {
		logrus.Fatalf("Error creating middlewares: %v", err)
	}

	d, err := daemon.NewDaemon(cli.Config, registryService, containerdRemote, pluginStore)
	if err != nil {
		return fmt.Errorf("Error starting daemon: %v", err)
	}

	d.StoreHosts(hosts)

	// validate after NewDaemon has restored enabled plugins. Dont change order.
	if err := validateAuthzPlugins(cli.Config.AuthorizationPlugins, pluginStore); err != nil {
		return fmt.Errorf("Error validating authorization plugin: %v", err)
	}

	if cli.Config.MetricsAddress != "" {
		if !d.HasExperimental() {
			return fmt.Errorf("metrics-addr is only supported when experimental is enabled")
		}
		if err := startMetricsServer(cli.Config.MetricsAddress); err != nil {
			return err
		}
	}

	name, _ := os.Hostname()

	// Use a buffered channel to pass changes from store watch API to daemon
	// A buffer allows store watch API and daemon processing to not wait for each other
	watchStream := make(chan *swarmapi.WatchMessage, 32)

	c, err := cluster.New(cluster.Config{
		Root:                   cli.Config.Root,
		Name:                   name,
		Backend:                d,
		NetworkSubnetsProvider: d,
		DefaultAdvertiseAddr:   cli.Config.SwarmDefaultAdvertiseAddr,
		RuntimeRoot:            cli.getSwarmRunRoot(),
		WatchStream:            watchStream,
	})
	if err != nil {
		logrus.Fatalf("Error creating cluster component: %v", err)
	}
	d.SetCluster(c)
	err = c.Start()
	if err != nil {
		logrus.Fatalf("Error starting cluster component: %v", err)
	}

	// Restart all autostart containers which has a swarm endpoint
	// and is not yet running now that we have successfully
	// initialized the cluster.
	d.RestartSwarmContainers()

	logrus.Info("Daemon has completed initialization")

	logrus.WithFields(logrus.Fields{
		"version":     dockerversion.Version,
		"commit":      dockerversion.GitCommit,
		"graphdriver": d.GraphDriverName(),
	}).Info("Docker daemon")

	cli.d = d

	initRouter(api, d, c)

	// process cluster change notifications
	watchCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go d.ProcessClusterNotifications(watchCtx, watchStream)

	cli.setupConfigReloadTrap()

	// The serve API routine never exits unless an error occurs
	// We need to start it as a goroutine and wait on it so
	// daemon doesn't exit
	serveAPIWait := make(chan error)
	go api.Wait(serveAPIWait)

	// after the daemon is done setting up we can notify systemd api
	notifySystem()

	// Daemon is fully initialized and handling API traffic
	// Wait for serve API to complete
	errAPI := <-serveAPIWait
	c.Cleanup()
	shutdownDaemon(d)
	containerdRemote.Cleanup()
	if errAPI != nil {
		return fmt.Errorf("Shutting down due to ServeAPI error: %v", errAPI)
	}

	return nil
}

func (cli *DaemonCli) reloadConfig() {
	reload := func(config *config.Config) {

		// Revalidate and reload the authorization plugins
		if err := validateAuthzPlugins(config.AuthorizationPlugins, cli.d.PluginStore); err != nil {
			logrus.Fatalf("Error validating authorization plugin: %v", err)
			return
		}
		cli.authzMiddleware.SetPlugins(config.AuthorizationPlugins)

		if err := cli.d.Reload(config); err != nil {
			logrus.Errorf("Error reconfiguring the daemon: %v", err)
			return
		}

		if config.IsValueSet("debug") {
			debugEnabled := debug.IsEnabled()
			switch {
			case debugEnabled && !config.Debug: // disable debug
				debug.Disable()
				cli.api.DisableProfiler()
			case config.Debug && !debugEnabled: // enable debug
				debug.Enable()
				cli.api.EnableProfiler()
			}

		}
	}

	if err := config.Reload(*cli.configFile, cli.flags, reload); err != nil {
		logrus.Error(err)
	}
}

func (cli *DaemonCli) stop() {
	cli.api.Close()
}

// shutdownDaemon just wraps daemon.Shutdown() to handle a timeout in case
// d.Shutdown() is waiting too long to kill container or worst it's
// blocked there
func shutdownDaemon(d *daemon.Daemon) {
	shutdownTimeout := d.ShutdownTimeout()
	ch := make(chan struct{})
	go func() {
		d.Shutdown()
		close(ch)
	}()
	if shutdownTimeout < 0 {
		<-ch
		logrus.Debug("Clean shutdown succeeded")
		return
	}
	select {
	case <-ch:
		logrus.Debug("Clean shutdown succeeded")
	case <-time.After(time.Duration(shutdownTimeout) * time.Second):
		logrus.Error("Force shutdown daemon")
	}
}

func loadDaemonCliConfig(opts daemonOptions) (*config.Config, error) {
	conf := opts.daemonConfig
	flags := opts.flags
	conf.Debug = opts.common.Debug
	conf.Hosts = opts.common.Hosts
	conf.LogLevel = opts.common.LogLevel
	conf.TLS = opts.common.TLS
	conf.TLSVerify = opts.common.TLSVerify
	conf.CommonTLSOptions = config.CommonTLSOptions{}

	if opts.common.TLSOptions != nil {
		conf.CommonTLSOptions.CAFile = opts.common.TLSOptions.CAFile
		conf.CommonTLSOptions.CertFile = opts.common.TLSOptions.CertFile
		conf.CommonTLSOptions.KeyFile = opts.common.TLSOptions.KeyFile
	}

	if conf.TrustKeyPath == "" {
		conf.TrustKeyPath = filepath.Join(
			getDaemonConfDir(conf.Root),
			defaultTrustKeyFile)
	}

	if flags.Changed("graph") && flags.Changed("data-root") {
		return nil, fmt.Errorf(`cannot specify both "--graph" and "--data-root" option`)
	}

	if opts.configFile != "" {
		c, err := config.MergeDaemonConfigurations(conf, flags, opts.configFile)
		if err != nil {
			if flags.Changed("config-file") || !os.IsNotExist(err) {
				return nil, fmt.Errorf("unable to configure the Docker daemon with file %s: %v\n", opts.configFile, err)
			}
		}
		// the merged configuration can be nil if the config file didn't exist.
		// leave the current configuration as it is if when that happens.
		if c != nil {
			conf = c
		}
	}

	if err := config.Validate(conf); err != nil {
		return nil, err
	}

	if flags.Changed("graph") {
		logrus.Warnf(`the "-g / --graph" flag is deprecated. Please use "--data-root" instead`)
	}

	// Labels of the docker engine used to allow multiple values associated with the same key.
	// This is deprecated in 1.13, and, be removed after 3 release cycles.
	// The following will check the conflict of labels, and report a warning for deprecation.
	//
	// TODO: After 3 release cycles (17.12) an error will be returned, and labels will be
	// sanitized to consolidate duplicate key-value pairs (config.Labels = newLabels):
	//
	// newLabels, err := daemon.GetConflictFreeLabels(config.Labels)
	// if err != nil {
	//	return nil, err
	// }
	// config.Labels = newLabels
	//
	if _, err := config.GetConflictFreeLabels(conf.Labels); err != nil {
		logrus.Warnf("Engine labels with duplicate keys and conflicting values have been deprecated: %s", err)
	}

	// Regardless of whether the user sets it to true or false, if they
	// specify TLSVerify at all then we need to turn on TLS
	if conf.IsValueSet(cliflags.FlagTLSVerify) {
		conf.TLS = true
	}

	// ensure that the log level is the one set after merging configurations
	cliflags.SetLogLevel(conf.LogLevel)

	return conf, nil
}

func initRouter(s *apiserver.Server, d *daemon.Daemon, c *cluster.Cluster) {
	decoder := runconfig.ContainerDecoder{}

	routers := []router.Router{
		// we need to add the checkpoint router before the container router or the DELETE gets masked
		checkpointrouter.NewRouter(d, decoder),
		container.NewRouter(d, decoder),
		image.NewRouter(d, decoder),
		systemrouter.NewRouter(d, c),
		volume.NewRouter(d),
		build.NewRouter(buildbackend.NewBackend(d, d), d),
		swarmrouter.NewRouter(c),
		pluginrouter.NewRouter(d.PluginManager()),
		distributionrouter.NewRouter(d),
	}

	if d.NetworkControllerEnabled() {
		routers = append(routers, network.NewRouter(d, c))
	}

	if d.HasExperimental() {
		for _, r := range routers {
			for _, route := range r.Routes() {
				if experimental, ok := route.(router.ExperimentalRoute); ok {
					experimental.Enable()
				}
			}
		}
	}

	s.InitRouter(debug.IsEnabled(), routers...)
}

func (cli *DaemonCli) initMiddlewares(s *apiserver.Server, cfg *apiserver.Config, pluginStore *plugin.Store) error {
	v := cfg.Version

	exp := middleware.NewExperimentalMiddleware(cli.Config.Experimental)
	s.UseMiddleware(exp)

	vm := middleware.NewVersionMiddleware(v, api.DefaultVersion, api.MinVersion)
	s.UseMiddleware(vm)

	if cfg.EnableCors || cfg.CorsHeaders != "" {
		c := middleware.NewCORSMiddleware(cfg.CorsHeaders)
		s.UseMiddleware(c)
	}

	cli.authzMiddleware = authorization.NewMiddleware(cli.Config.AuthorizationPlugins, pluginStore)
	cli.Config.AuthzMiddleware = cli.authzMiddleware
	s.UseMiddleware(cli.authzMiddleware)
	return nil
}

// validates that the plugins requested with the --authorization-plugin flag are valid AuthzDriver
// plugins present on the host and available to the daemon
func validateAuthzPlugins(requestedPlugins []string, pg plugingetter.PluginGetter) error {
	for _, reqPlugin := range requestedPlugins {
		if _, err := pg.Get(reqPlugin, authorization.AuthZApiImplements, plugingetter.Lookup); err != nil {
			return err
		}
	}
	return nil
}