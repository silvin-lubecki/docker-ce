package service // import "github.com/docker/docker/integration/service"

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/integration/internal/container"
	"github.com/docker/docker/integration/internal/swarm"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
)

func TestDockerNetworkConnectAlias(t *testing.T) {
	defer setupTest(t)()
	d := swarm.NewSwarm(t, testEnv)
	defer d.Stop(t)
	client := d.NewClientT(t)
	defer client.Close()
	ctx := context.Background()

	name := "test-alias"
	_, err := client.NetworkCreate(ctx, name, types.NetworkCreate{
		Driver:     "overlay",
		Attachable: true,
	})
	assert.NilError(t, err)

	container.Create(t, ctx, client, container.WithName("ng1"), func(c *container.TestContainerConfig) {
		c.NetworkingConfig = &network.NetworkingConfig{
			map[string]*network.EndpointSettings{
				name: {},
			},
		}
	})

	err = client.NetworkConnect(ctx, name, "ng1", &network.EndpointSettings{
		Aliases: []string{
			"aaa",
		},
	})
	assert.NilError(t, err)

	err = client.ContainerStart(ctx, "ng1", types.ContainerStartOptions{})
	assert.NilError(t, err)

	ng1, err := client.ContainerInspect(ctx, "ng1")
	assert.NilError(t, err)
	assert.Check(t, is.Equal(len(ng1.NetworkSettings.Networks[name].Aliases), 2))
	assert.Check(t, is.Equal(ng1.NetworkSettings.Networks[name].Aliases[0], "aaa"))

	container.Create(t, ctx, client, container.WithName("ng2"), func(c *container.TestContainerConfig) {
		c.NetworkingConfig = &network.NetworkingConfig{
			map[string]*network.EndpointSettings{
				name: {},
			},
		}
	})

	err = client.NetworkConnect(ctx, name, "ng2", &network.EndpointSettings{
		Aliases: []string{
			"bbb",
		},
	})
	assert.NilError(t, err)

	err = client.ContainerStart(ctx, "ng2", types.ContainerStartOptions{})
	assert.NilError(t, err)

	ng2, err := client.ContainerInspect(ctx, "ng2")
	assert.NilError(t, err)
	assert.Check(t, is.Equal(len(ng2.NetworkSettings.Networks[name].Aliases), 2))
	assert.Check(t, is.Equal(ng2.NetworkSettings.Networks[name].Aliases[0], "bbb"))
}
