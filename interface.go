package client

import (
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/swarm"
	volumetypes "github.com/docker/docker/api/types/volume"
	"golang.org/x/net/context"
)

// CommonAPIClient is the common methods between stable and experimental versions of APIClient.
type CommonAPIClient interface {
	ContainerAPIClient
	ImageAPIClient
	NodeAPIClient
	NetworkAPIClient
	ServiceAPIClient
	SwarmAPIClient
	SystemAPIClient
	VolumeAPIClient
	ClientVersion() string
	ServerVersion(ctx context.Context) (types.Version, error)
	UpdateClientVersion(v string)
}

// ContainerAPIClient defines API client methods for the containers
type ContainerAPIClient interface {
	ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
	ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
	ContainerDiff(ctx context.Context, container string) ([]types.ContainerChange, error)
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
	ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
	ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
	ContainerExport(ctx context.Context, container string) (io.ReadCloser, error)
	ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error)
	ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error)
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
	ContainerPause(ctx context.Context, container string) error
	ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error
	ContainerRename(ctx context.Context, container, newContainerName string) error
	ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error
	ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error
	ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error)
	ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error)
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerStop(ctx context.Context, container string, timeout *time.Duration) error
	ContainerTop(ctx context.Context, container string, arguments []string) (types.ContainerProcessList, error)
	ContainerUnpause(ctx context.Context, container string) error
	ContainerUpdate(ctx context.Context, container string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
	ContainerWait(ctx context.Context, container string) (int, error)
	CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
	CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
	ContainersPrune(ctx context.Context, cfg types.ContainersPruneConfig) (types.ContainersPruneReport, error)
}

// ImageAPIClient defines API client methods for the images
type ImageAPIClient interface {
	ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
	ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
	ImageHistory(ctx context.Context, image string) ([]types.ImageHistory, error)
	ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
	ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error)
	ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
	ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
	ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
	ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error)
	ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDelete, error)
	ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
	ImageSave(ctx context.Context, images []string) (io.ReadCloser, error)
	ImageTag(ctx context.Context, image, ref string) error
	ImagesPrune(ctx context.Context, cfg types.ImagesPruneConfig) (types.ImagesPruneReport, error)
}

// NetworkAPIClient defines API client methods for the networks
type NetworkAPIClient interface {
	NetworkConnect(ctx context.Context, networkID, container string, config *network.EndpointSettings) error
	NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
	NetworkDisconnect(ctx context.Context, networkID, container string, force bool) error
	NetworkInspect(ctx context.Context, networkID string) (types.NetworkResource, error)
	NetworkInspectWithRaw(ctx context.Context, networkID string) (types.NetworkResource, []byte, error)
	NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
	NetworkRemove(ctx context.Context, networkID string) error
	NetworksPrune(ctx context.Context, cfg types.NetworksPruneConfig) (types.NetworksPruneReport, error)
}

// NodeAPIClient defines API client methods for the nodes
type NodeAPIClient interface {
	NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
	NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
	NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
	NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
}

// ServiceAPIClient defines API client methods for the services
type ServiceAPIClient interface {
	ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
	ServiceInspectWithRaw(ctx context.Context, serviceID string) (swarm.Service, []byte, error)
	ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
	ServiceRemove(ctx context.Context, serviceID string) error
	ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) error
	TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
	TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
}

// SwarmAPIClient defines API client methods for the swarm
type SwarmAPIClient interface {
	SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
	SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
	SwarmLeave(ctx context.Context, force bool) error
	SwarmInspect(ctx context.Context) (swarm.Swarm, error)
	SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
}

// SystemAPIClient defines API client methods for the system
type SystemAPIClient interface {
	Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
	Info(ctx context.Context) (types.Info, error)
	RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
	DiskUsage(ctx context.Context) (types.DiskUsage, error)
	Ping(ctx context.Context) (bool, error)
}

// VolumeAPIClient defines API client methods for the volumes
type VolumeAPIClient interface {
	VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
	VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
	VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
	VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
	VolumeRemove(ctx context.Context, volumeID string, force bool) error
	VolumesPrune(ctx context.Context, cfg types.VolumesPruneConfig) (types.VolumesPruneReport, error)
}
