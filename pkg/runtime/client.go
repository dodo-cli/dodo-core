package runtime

import (
	"context"
	"net"
	"time"

	"github.com/dodo-cli/dodo-core/pkg/appconfig"
	"google.golang.org/grpc"
	"k8s.io/cri-api/pkg/apis"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const connectionTimout = 15 * time.Second

var _ Client = &client{}

type Client interface {
	cri.RuntimeService
	cri.ImageManagerService
}

type client struct {
	timeout       time.Duration
	runtimeClient runtimeapi.RuntimeServiceClient
	imageClient   runtimeapi.ImageServiceClient
}

func NewClient() (Client, error) {
	conn, err := grpc.Dial(
		appconfig.GetCRIEndpoint(),
		grpc.WithInsecure(),
		grpc.WithTimeout(connectionTimout),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		timeout:       connectionTimout,
		runtimeClient: runtimeapi.NewRuntimeServiceClient(conn),
		imageClient:   runtimeapi.NewImageServiceClient(conn),
	}, nil
}

func (c *client) Version(apiVersion string) (*runtimeapi.VersionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.runtimeClient.Version(ctx, &runtimeapi.VersionRequest{Version: apiVersion})
}

func (c *client) RunPodSandbox(config *runtimeapi.PodSandboxConfig, runtimeHandler string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout*2)
	defer cancel()

	resp, err := c.runtimeClient.RunPodSandbox(ctx, &runtimeapi.RunPodSandboxRequest{
		Config:         config,
		RuntimeHandler: runtimeHandler,
	})
	if err != nil {
		return "", err
	}

	return resp.PodSandboxId, nil
}

func (c *client) StopPodSandbox(podSandBoxID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.StopPodSandbox(ctx, &runtimeapi.StopPodSandboxRequest{PodSandboxId: podSandBoxID})

	return err
}

func (c *client) RemovePodSandbox(podSandBoxID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.RemovePodSandbox(ctx, &runtimeapi.RemovePodSandboxRequest{PodSandboxId: podSandBoxID})

	return err
}

func (c *client) PodSandboxStatus(podSandBoxID string) (*runtimeapi.PodSandboxStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.PodSandboxStatus(ctx, &runtimeapi.PodSandboxStatusRequest{PodSandboxId: podSandBoxID})
	if err != nil {
		return nil, err
	}

	return resp.Status, nil
}

func (c *client) ListPodSandbox(filter *runtimeapi.PodSandboxFilter) ([]*runtimeapi.PodSandbox, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.ListPodSandbox(ctx, &runtimeapi.ListPodSandboxRequest{Filter: filter})
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}

func (c *client) CreateContainer(podSandBoxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.CreateContainer(ctx, &runtimeapi.CreateContainerRequest{
		PodSandboxId:  podSandBoxID,
		Config:        config,
		SandboxConfig: sandboxConfig,
	})
	if err != nil {
		return "", err
	}

	return resp.ContainerId, nil
}

func (c *client) StartContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.StartContainer(ctx, &runtimeapi.StartContainerRequest{ContainerId: containerID})

	return err
}

func (c *client) StopContainer(containerID string, timeout int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	_, err := c.runtimeClient.StopContainer(ctx, &runtimeapi.StopContainerRequest{
		ContainerId: containerID,
		Timeout:     timeout,
	})

	return err
}

func (c *client) RemoveContainer(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.RemoveContainer(ctx, &runtimeapi.RemoveContainerRequest{ContainerId: containerID})

	return err
}

func (c *client) ListContainers(filter *runtimeapi.ContainerFilter) ([]*runtimeapi.Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.ListContainers(ctx, &runtimeapi.ListContainersRequest{Filter: filter})
	if err != nil {
		return nil, err
	}

	return resp.Containers, nil
}

func (c *client) ContainerStatus(containerID string) (*runtimeapi.ContainerStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.ContainerStatus(ctx, &runtimeapi.ContainerStatusRequest{ContainerId: containerID})
	if err != nil {
		return nil, err
	}

	return resp.Status, nil
}

func (c *client) UpdateContainerResources(containerID string, resources *runtimeapi.LinuxContainerResources) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.UpdateContainerResources(ctx, &runtimeapi.UpdateContainerResourcesRequest{
		ContainerId: containerID,
		Linux:       resources,
	})

	return err
}

func (c *client) ExecSync(containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error) {
	var ctx context.Context
	var cancel context.CancelFunc

	if timeout != 0 {
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout+timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}

	defer cancel()

	resp, err := c.runtimeClient.ExecSync(ctx, &runtimeapi.ExecSyncRequest{
		ContainerId: containerID,
		Cmd:         cmd,
		Timeout:     int64(timeout.Seconds()),
	})
	if err != nil {
		return nil, nil, err
	}

	return resp.Stdout, resp.Stderr, err
}

func (c *client) Exec(req *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.runtimeClient.Exec(ctx, req)
}

func (c *client) Attach(req *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.runtimeClient.Attach(ctx, req)
}

func (c *client) PortForward(req *runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.runtimeClient.PortForward(ctx, req)
}

func (c *client) UpdateRuntimeConfig(runtimeConfig *runtimeapi.RuntimeConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.UpdateRuntimeConfig(ctx, &runtimeapi.UpdateRuntimeConfigRequest{
		RuntimeConfig: runtimeConfig,
	})

	return err
}

func (c *client) Status() (*runtimeapi.RuntimeStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.Status(ctx, &runtimeapi.StatusRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Status, nil
}

func (c *client) ContainerStats(containerID string) (*runtimeapi.ContainerStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.runtimeClient.ContainerStats(ctx, &runtimeapi.ContainerStatsRequest{ContainerId: containerID})
	if err != nil {
		return nil, err
	}

	return resp.Stats, nil
}

func (c *client) ListContainerStats(filter *runtimeapi.ContainerStatsFilter) ([]*runtimeapi.ContainerStats, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.runtimeClient.ListContainerStats(ctx, &runtimeapi.ListContainerStatsRequest{Filter: filter})
	if err != nil {
		return nil, err
	}

	return resp.Stats, nil
}

func (c *client) ReopenContainerLog(containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.runtimeClient.ReopenContainerLog(ctx, &runtimeapi.ReopenContainerLogRequest{ContainerId: containerID})

	return err
}

func (c *client) ListImages(filter *runtimeapi.ImageFilter) ([]*runtimeapi.Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.imageClient.ListImages(ctx, &runtimeapi.ListImagesRequest{Filter: filter})
	if err != nil {
		return nil, err
	}

	return resp.Images, nil
}

func (c *client) ImageStatus(image *runtimeapi.ImageSpec) (*runtimeapi.Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.imageClient.ImageStatus(ctx, &runtimeapi.ImageStatusRequest{Image: image})
	if err != nil {
		return nil, err
	}

	return resp.Image, nil
}

func (c *client) PullImage(image *runtimeapi.ImageSpec, auth *runtimeapi.AuthConfig, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.imageClient.PullImage(ctx, &runtimeapi.PullImageRequest{
		Image:         image,
		Auth:          auth,
		SandboxConfig: podSandboxConfig,
	})
	if err != nil {
		return "", err
	}

	return resp.ImageRef, nil
}

func (c *client) RemoveImage(image *runtimeapi.ImageSpec) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.imageClient.RemoveImage(ctx, &runtimeapi.RemoveImageRequest{Image: image})

	return err
}

func (c *client) ImageFsInfo() ([]*runtimeapi.FilesystemUsage, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.imageClient.ImageFsInfo(ctx, &runtimeapi.ImageFsInfoRequest{})
	if err != nil {
		return nil, err
	}

	return resp.ImageFilesystems, nil
}
