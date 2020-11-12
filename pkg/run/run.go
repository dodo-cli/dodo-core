package run

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/configuration"
	"github.com/dodo-cli/dodo-core/pkg/runtime"
	"github.com/dodo-cli/dodo-core/pkg/types"
	log "github.com/hashicorp/go-hclog"
	"github.com/moby/term"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func RunContainer(overrides *types.Backdrop) error {
	config := GetConfig(overrides)

	if len(config.ContainerName) == 0 {
		id := make([]byte, 8)
		if _, err := rand.Read(id); err != nil {
			panic(err)
		}

		config.ContainerName = fmt.Sprintf("%s-%s", config.Name, hex.EncodeToString(id))
	}

	rt, err := runtime.NewClient()
	if err != nil {
		return err
	}

	imageID, err := rt.PullImage(
		&runtimeapi.ImageSpec{Image: config.ImageId},
		&runtimeapi.AuthConfig{},
		&runtimeapi.PodSandboxConfig{},
	)
	if err != nil {
		return err
	}

	config.ImageId = imageID

	randomID := make([]byte, 20)
	if _, err := rand.Read(randomID); err != nil {
		return err
	}

	tty := term.IsTerminal(os.Stdin.Fd()) && term.IsTerminal(os.Stdout.Fd())
	tmpPath := fmt.Sprintf("/tmp/dodo-%s/", hex.EncodeToString(randomID))
	sandboxConfig, containerConfig := config.RuntimeConfig(tmpPath, tty, true)

	sandboxID, err := rt.RunPodSandbox(sandboxConfig, "")
	if err != nil {
		return err
	}

	containerID, err := rt.CreateContainer(sandboxID, containerConfig, sandboxConfig)
	if err != nil {
		return err
	}

	if len(config.Entrypoint.Script) > 0 {
		encoded := base64.StdEncoding.EncodeToString([]byte(config.Entrypoint.Script + "\n"))
		script := fmt.Sprintf("echo %s | base64 -d > %s", encoded, path.Join(tmpPath, "entrypoint"))

		_, stderr, err := rt.ExecSync(containerID, []string{"/bin/sh", "-c", script}, 5*time.Second)
		if err != nil {
			return fmt.Errorf("%s: %w", stderr, err)
		}
	}

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		err := p.(configuration.Configuration).Provision(containerID)
		if err != nil {
			log.Default().Warn("could not provision", "error", err)
		}
	}

	resp, err := rt.Attach(&runtimeapi.AttachRequest{
		ContainerId: containerID,
		Tty:         tty,
		Stdin:       true,
		Stdout:      true,
		Stderr:      true,
	})
	if err != nil {
		return err
	}

	attachURL, err := url.Parse(resp.Url)
	if err != nil {
		return err
	}

	executor, err := remotecommand.NewSPDYExecutor(
		&rest.Config{},
		"POST",
		attachURL,
	)
	if err != nil {
		return err
	}

	opts := remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    tty,
	}

	fd := os.Stdin.Fd()
	if term.IsTerminal(fd) {
		resizeChannel := make(chan *remotecommand.TerminalSize)
		winch := make(chan os.Signal, 1)
		signal.Notify(winch, syscall.SIGWINCH)

		go func() {
			for range winch {
				ws, err := term.GetWinsize(fd)
				if err != nil {
					continue
				}

				if ws.Height == 0 && ws.Width == 0 {
					continue
				}

				resizeChannel <- &remotecommand.TerminalSize{Height: ws.Height, Width: ws.Width}
			}
		}()

		opts.TerminalSizeQueue = &resizer{resizeChannel: resizeChannel}
	}

	return executor.Stream(opts)
}

func GetConfig(overrides *types.Backdrop) *types.Backdrop {
	config := &types.Backdrop{Name: overrides.Name, Entrypoint: &types.Entrypoint{}}

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		conf, err := p.(configuration.Configuration).UpdateConfiguration(config)
		if err != nil {
			log.L().Warn("could not get config", "error", err)
			continue
		}

		config.Merge(conf)
	}

	config.Merge(overrides)
	log.L().Debug("assembled configuration", "backdrop", config)
	return config
}

type resizer struct {
	resizeChannel chan *remotecommand.TerminalSize
}

func (r *resizer) Next() *remotecommand.TerminalSize {
	return <-r.resizeChannel
}
