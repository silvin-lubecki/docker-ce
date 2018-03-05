package swarm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/docker/cli/internal/test"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
	"github.com/gotestyourself/gotestyourself/golden"
	"github.com/pkg/errors"
)

func TestSwarmInitErrorOnAPIFailure(t *testing.T) {
	testCases := []struct {
		name                  string
		flags                 map[string]string
		swarmInitFunc         func() (string, error)
		swarmInspectFunc      func() (swarm.Swarm, error)
		swarmGetUnlockKeyFunc func() (types.SwarmUnlockKeyResponse, error)
		nodeInspectFunc       func() (swarm.Node, []byte, error)
		expectedError         string
	}{
		{
			name: "init-failed",
			swarmInitFunc: func() (string, error) {
				return "", errors.Errorf("error initializing the swarm")
			},
			expectedError: "error initializing the swarm",
		},
		{
			name: "init-failed-with-ip-choice",
			swarmInitFunc: func() (string, error) {
				return "", errors.Errorf("could not choose an IP address to advertise")
			},
			expectedError: "could not choose an IP address to advertise - specify one with --advertise-addr",
		},
		{
			name: "swarm-inspect-after-init-failed",
			swarmInspectFunc: func() (swarm.Swarm, error) {
				return swarm.Swarm{}, errors.Errorf("error inspecting the swarm")
			},
			expectedError: "error inspecting the swarm",
		},
		{
			name: "node-inspect-after-init-failed",
			nodeInspectFunc: func() (swarm.Node, []byte, error) {
				return swarm.Node{}, []byte{}, errors.Errorf("error inspecting the node")
			},
			expectedError: "error inspecting the node",
		},
		{
			name: "swarm-get-unlock-key-after-init-failed",
			flags: map[string]string{
				flagAutolock: "true",
			},
			swarmGetUnlockKeyFunc: func() (types.SwarmUnlockKeyResponse, error) {
				return types.SwarmUnlockKeyResponse{}, errors.Errorf("error getting swarm unlock key")
			},
			expectedError: "could not fetch unlock key: error getting swarm unlock key",
		},
	}
	for _, tc := range testCases {
		cmd := newInitCommand(
			test.NewFakeCli(&fakeClient{
				swarmInitFunc:         tc.swarmInitFunc,
				swarmInspectFunc:      tc.swarmInspectFunc,
				swarmGetUnlockKeyFunc: tc.swarmGetUnlockKeyFunc,
				nodeInspectFunc:       tc.nodeInspectFunc,
			}))
		for key, value := range tc.flags {
			cmd.Flags().Set(key, value)
		}
		cmd.SetOutput(ioutil.Discard)
		assert.Check(t, is.Error(cmd.Execute(), tc.expectedError))
	}
}

func TestSwarmInit(t *testing.T) {
	testCases := []struct {
		name                  string
		flags                 map[string]string
		swarmInitFunc         func() (string, error)
		swarmInspectFunc      func() (swarm.Swarm, error)
		swarmGetUnlockKeyFunc func() (types.SwarmUnlockKeyResponse, error)
		nodeInspectFunc       func() (swarm.Node, []byte, error)
	}{
		{
			name: "init",
			swarmInitFunc: func() (string, error) {
				return "nodeID", nil
			},
		},
		{
			name: "init-autolock",
			flags: map[string]string{
				flagAutolock: "true",
			},
			swarmInitFunc: func() (string, error) {
				return "nodeID", nil
			},
			swarmGetUnlockKeyFunc: func() (types.SwarmUnlockKeyResponse, error) {
				return types.SwarmUnlockKeyResponse{
					UnlockKey: "unlock-key",
				}, nil
			},
		},
	}
	for _, tc := range testCases {
		cli := test.NewFakeCli(&fakeClient{
			swarmInitFunc:         tc.swarmInitFunc,
			swarmInspectFunc:      tc.swarmInspectFunc,
			swarmGetUnlockKeyFunc: tc.swarmGetUnlockKeyFunc,
			nodeInspectFunc:       tc.nodeInspectFunc,
		})
		cmd := newInitCommand(cli)
		for key, value := range tc.flags {
			cmd.Flags().Set(key, value)
		}
		assert.Check(t, cmd.Execute())
		golden.Assert(t, cli.OutBuffer().String(), fmt.Sprintf("init-%s.golden", tc.name))
	}
}
