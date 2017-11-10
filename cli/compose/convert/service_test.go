package convert

import (
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	composetypes "github.com/docker/cli/cli/compose/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/swarm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertRestartPolicyFromNone(t *testing.T) {
	policy, err := convertRestartPolicy("no", nil)
	assert.NoError(t, err)
	assert.Equal(t, (*swarm.RestartPolicy)(nil), policy)
}

func TestConvertRestartPolicyFromUnknown(t *testing.T) {
	_, err := convertRestartPolicy("unknown", nil)
	assert.EqualError(t, err, "unknown restart policy: unknown")
}

func TestConvertRestartPolicyFromAlways(t *testing.T) {
	policy, err := convertRestartPolicy("always", nil)
	expected := &swarm.RestartPolicy{
		Condition: swarm.RestartPolicyConditionAny,
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, policy)
}

func TestConvertRestartPolicyFromFailure(t *testing.T) {
	policy, err := convertRestartPolicy("on-failure:4", nil)
	attempts := uint64(4)
	expected := &swarm.RestartPolicy{
		Condition:   swarm.RestartPolicyConditionOnFailure,
		MaxAttempts: &attempts,
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, policy)
}

func strPtr(val string) *string {
	return &val
}

func TestConvertEnvironment(t *testing.T) {
	source := map[string]*string{
		"foo": strPtr("bar"),
		"key": strPtr("value"),
	}
	env := convertEnvironment(source)
	sort.Strings(env)
	assert.Equal(t, []string{"foo=bar", "key=value"}, env)
}

func TestConvertExtraHosts(t *testing.T) {
	source := composetypes.HostsList{
		"zulu:127.0.0.2",
		"alpha:127.0.0.1",
		"zulu:ff02::1",
	}
	assert.Equal(t, []string{"127.0.0.2 zulu", "127.0.0.1 alpha", "ff02::1 zulu"}, convertExtraHosts(source))
}

func TestConvertResourcesFull(t *testing.T) {
	source := composetypes.Resources{
		Limits: &composetypes.Resource{
			NanoCPUs:    "0.003",
			MemoryBytes: composetypes.UnitBytes(300000000),
		},
		Reservations: &composetypes.Resource{
			NanoCPUs:    "0.002",
			MemoryBytes: composetypes.UnitBytes(200000000),
		},
	}
	resources, err := convertResources(source)
	assert.NoError(t, err)

	expected := &swarm.ResourceRequirements{
		Limits: &swarm.Resources{
			NanoCPUs:    3000000,
			MemoryBytes: 300000000,
		},
		Reservations: &swarm.Resources{
			NanoCPUs:    2000000,
			MemoryBytes: 200000000,
		},
	}
	assert.Equal(t, expected, resources)
}

func TestConvertResourcesOnlyMemory(t *testing.T) {
	source := composetypes.Resources{
		Limits: &composetypes.Resource{
			MemoryBytes: composetypes.UnitBytes(300000000),
		},
		Reservations: &composetypes.Resource{
			MemoryBytes: composetypes.UnitBytes(200000000),
		},
	}
	resources, err := convertResources(source)
	assert.NoError(t, err)

	expected := &swarm.ResourceRequirements{
		Limits: &swarm.Resources{
			MemoryBytes: 300000000,
		},
		Reservations: &swarm.Resources{
			MemoryBytes: 200000000,
		},
	}
	assert.Equal(t, expected, resources)
}

func TestConvertHealthcheck(t *testing.T) {
	retries := uint64(10)
	timeout := 30 * time.Second
	interval := 2 * time.Millisecond
	source := &composetypes.HealthCheckConfig{
		Test:     []string{"EXEC", "touch", "/foo"},
		Timeout:  &timeout,
		Interval: &interval,
		Retries:  &retries,
	}
	expected := &container.HealthConfig{
		Test:     source.Test,
		Timeout:  timeout,
		Interval: interval,
		Retries:  10,
	}

	healthcheck, err := convertHealthcheck(source)
	assert.NoError(t, err)
	assert.Equal(t, expected, healthcheck)
}

func TestConvertHealthcheckDisable(t *testing.T) {
	source := &composetypes.HealthCheckConfig{Disable: true}
	expected := &container.HealthConfig{
		Test: []string{"NONE"},
	}

	healthcheck, err := convertHealthcheck(source)
	assert.NoError(t, err)
	assert.Equal(t, expected, healthcheck)
}

func TestConvertHealthcheckDisableWithTest(t *testing.T) {
	source := &composetypes.HealthCheckConfig{
		Disable: true,
		Test:    []string{"EXEC", "touch"},
	}
	_, err := convertHealthcheck(source)
	assert.EqualError(t, err, "test and disable can't be set at the same time")
}

func TestConvertEndpointSpec(t *testing.T) {
	source := []composetypes.ServicePortConfig{
		{
			Protocol:  "udp",
			Target:    53,
			Published: 1053,
			Mode:      "host",
		},
		{
			Target:    8080,
			Published: 80,
		},
	}
	endpoint, err := convertEndpointSpec("vip", source)

	expected := swarm.EndpointSpec{
		Mode: swarm.ResolutionMode(strings.ToLower("vip")),
		Ports: []swarm.PortConfig{
			{
				TargetPort:    8080,
				PublishedPort: 80,
			},
			{
				Protocol:      "udp",
				TargetPort:    53,
				PublishedPort: 1053,
				PublishMode:   "host",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, *endpoint)
}

func TestConvertServiceNetworksOnlyDefault(t *testing.T) {
	networkConfigs := networkMap{}

	configs, err := convertServiceNetworks(
		nil, networkConfigs, NewNamespace("foo"), "service")

	expected := []swarm.NetworkAttachmentConfig{
		{
			Target:  "foo_default",
			Aliases: []string{"service"},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, configs)
}

func TestConvertServiceNetworks(t *testing.T) {
	networkConfigs := networkMap{
		"front": composetypes.NetworkConfig{
			External: composetypes.External{
				External: true,
				Name:     "fronttier",
			},
		},
		"back": composetypes.NetworkConfig{},
	}
	networks := map[string]*composetypes.ServiceNetworkConfig{
		"front": {
			Aliases: []string{"something"},
		},
		"back": {
			Aliases: []string{"other"},
		},
	}

	configs, err := convertServiceNetworks(
		networks, networkConfigs, NewNamespace("foo"), "service")

	expected := []swarm.NetworkAttachmentConfig{
		{
			Target:  "foo_back",
			Aliases: []string{"other", "service"},
		},
		{
			Target:  "fronttier",
			Aliases: []string{"something", "service"},
		},
	}

	sortedConfigs := byTargetSort(configs)
	sort.Sort(&sortedConfigs)

	assert.NoError(t, err)
	assert.Equal(t, expected, []swarm.NetworkAttachmentConfig(sortedConfigs))
}

func TestConvertServiceNetworksCustomDefault(t *testing.T) {
	networkConfigs := networkMap{
		"default": composetypes.NetworkConfig{
			External: composetypes.External{
				External: true,
				Name:     "custom",
			},
		},
	}
	networks := map[string]*composetypes.ServiceNetworkConfig{}

	configs, err := convertServiceNetworks(
		networks, networkConfigs, NewNamespace("foo"), "service")

	expected := []swarm.NetworkAttachmentConfig{
		{
			Target:  "custom",
			Aliases: []string{"service"},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, []swarm.NetworkAttachmentConfig(configs))
}

type byTargetSort []swarm.NetworkAttachmentConfig

func (s byTargetSort) Len() int {
	return len(s)
}

func (s byTargetSort) Less(i, j int) bool {
	return strings.Compare(s[i].Target, s[j].Target) < 0
}

func (s byTargetSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestConvertDNSConfigEmpty(t *testing.T) {
	dnsConfig, err := convertDNSConfig(nil, nil)

	assert.NoError(t, err)
	assert.Equal(t, (*swarm.DNSConfig)(nil), dnsConfig)
}

var (
	nameservers = []string{"8.8.8.8", "9.9.9.9"}
	search      = []string{"dc1.example.com", "dc2.example.com"}
)

func TestConvertDNSConfigAll(t *testing.T) {
	dnsConfig, err := convertDNSConfig(nameservers, search)
	assert.NoError(t, err)
	assert.Equal(t, &swarm.DNSConfig{
		Nameservers: nameservers,
		Search:      search,
	}, dnsConfig)
}

func TestConvertDNSConfigNameservers(t *testing.T) {
	dnsConfig, err := convertDNSConfig(nameservers, nil)
	assert.NoError(t, err)
	assert.Equal(t, &swarm.DNSConfig{
		Nameservers: nameservers,
		Search:      nil,
	}, dnsConfig)
}

func TestConvertDNSConfigSearch(t *testing.T) {
	dnsConfig, err := convertDNSConfig(nil, search)
	assert.NoError(t, err)
	assert.Equal(t, &swarm.DNSConfig{
		Nameservers: nil,
		Search:      search,
	}, dnsConfig)
}

func TestConvertCredentialSpec(t *testing.T) {
	swarmSpec, err := convertCredentialSpec(composetypes.CredentialSpecConfig{})
	assert.NoError(t, err)
	assert.Nil(t, swarmSpec)

	swarmSpec, err = convertCredentialSpec(composetypes.CredentialSpecConfig{
		File: "/foo",
	})
	assert.NoError(t, err)
	assert.Equal(t, swarmSpec.File, "/foo")
	assert.Equal(t, swarmSpec.Registry, "")

	swarmSpec, err = convertCredentialSpec(composetypes.CredentialSpecConfig{
		Registry: "foo",
	})
	assert.NoError(t, err)
	assert.Equal(t, swarmSpec.File, "")
	assert.Equal(t, swarmSpec.Registry, "foo")

	swarmSpec, err = convertCredentialSpec(composetypes.CredentialSpecConfig{
		File:     "/asdf",
		Registry: "foo",
	})
	assert.Error(t, err)
	assert.Nil(t, swarmSpec)
}

func TestConvertUpdateConfigOrder(t *testing.T) {
	// test default behavior
	updateConfig := convertUpdateConfig(&composetypes.UpdateConfig{})
	assert.Equal(t, "", updateConfig.Order)

	// test start-first
	updateConfig = convertUpdateConfig(&composetypes.UpdateConfig{
		Order: "start-first",
	})
	assert.Equal(t, updateConfig.Order, "start-first")

	// test stop-first
	updateConfig = convertUpdateConfig(&composetypes.UpdateConfig{
		Order: "stop-first",
	})
	assert.Equal(t, updateConfig.Order, "stop-first")
}

func TestConvertFileObject(t *testing.T) {
	namespace := NewNamespace("testing")
	config := composetypes.FileReferenceConfig{
		Source: "source",
		Target: "target",
		UID:    "user",
		GID:    "group",
		Mode:   uint32Ptr(0644),
	}
	swarmRef, err := convertFileObject(namespace, config, lookupConfig)
	require.NoError(t, err)

	expected := swarmReferenceObject{
		Name: "testing_source",
		File: swarmReferenceTarget{
			Name: config.Target,
			UID:  config.UID,
			GID:  config.GID,
			Mode: os.FileMode(0644),
		},
	}
	assert.Equal(t, expected, swarmRef)
}

func lookupConfig(key string) (composetypes.FileObjectConfig, error) {
	if key != "source" {
		return composetypes.FileObjectConfig{}, errors.New("bad key")
	}
	return composetypes.FileObjectConfig{}, nil
}

func TestConvertFileObjectDefaults(t *testing.T) {
	namespace := NewNamespace("testing")
	config := composetypes.FileReferenceConfig{Source: "source"}
	swarmRef, err := convertFileObject(namespace, config, lookupConfig)
	require.NoError(t, err)

	expected := swarmReferenceObject{
		Name: "testing_source",
		File: swarmReferenceTarget{
			Name: config.Source,
			UID:  "0",
			GID:  "0",
			Mode: os.FileMode(0444),
		},
	}
	assert.Equal(t, expected, swarmRef)
}
