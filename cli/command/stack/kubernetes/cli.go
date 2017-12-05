package kubernetes

import (
	"os"
	"path/filepath"

	"github.com/docker/cli/cli/command"
	composev1beta1 "github.com/docker/cli/kubernetes/client/clientset_generated/clientset/typed/compose/v1beta1"
	"github.com/docker/docker/pkg/homedir"
	"github.com/spf13/cobra"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeCli holds kubernetes specifics (client, namespace) with the command.Cli
type KubeCli struct {
	command.Cli
	kubeConfig    *restclient.Config
	kubeNamespace string
}

// WrapCli wraps command.Cli with kubernetes specifics
func WrapCli(dockerCli command.Cli, cmd *cobra.Command) (*KubeCli, error) {
	var err error
	cli := &KubeCli{
		Cli:           dockerCli,
		kubeNamespace: "default",
	}
	if cmd.PersistentFlags().Changed("namespace") {
		cli.kubeNamespace, err = cmd.PersistentFlags().GetString("namespace")
		if err != nil {
			return nil, err
		}
	}
	kubeConfig := ""
	if cmd.PersistentFlags().Changed("kubeconfig") {
		kubeConfig, err = cmd.PersistentFlags().GetString("kubeconfig")
		if err != nil {
			return nil, err
		}
	}
	if kubeConfig == "" {
		if config := os.Getenv("KUBECONFIG"); config != "" {
			kubeConfig = config
		} else {
			kubeConfig = filepath.Join(homedir.Get(), ".kube/config")
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	cli.kubeConfig = config

	return cli, nil
}

func (c *KubeCli) composeClient() (*Factory, error) {
	return NewFactory(c.kubeNamespace, c.kubeConfig)
}

func (c *KubeCli) stacks() (composev1beta1.StackInterface, error) {
	err := APIPresent(c.kubeConfig)
	if err != nil {
		return nil, err
	}

	clientSet, err := composev1beta1.NewForConfig(c.kubeConfig)
	if err != nil {
		return nil, err
	}

	return clientSet.Stacks(c.kubeNamespace), nil
}
