package imageplugin

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/garden-shed/rootfs_provider"
	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry/gunk/command_runner"
	errorwrapper "github.com/pkg/errors"
	"github.com/tscolari/lagregator"
)

//go:generate counterfeiter . CommandCreator
type CommandCreator interface {
	CreateCommand(log lager.Logger, handle string, spec rootfs_provider.Spec) (*exec.Cmd, error)
	DestroyCommand(log lager.Logger, handle string) (*exec.Cmd, error)
	MetricsCommand(log lager.Logger, handle string) (*exec.Cmd, error)
	GCCommand(log lager.Logger) (*exec.Cmd, error)
}

type CompositeImagePlugin struct {
	unprivilegedCommandCreator CommandCreator
	commandRunner              command_runner.CommandRunner
	defaultRootfs              string
}

func New(unprivilegedCommandCreator CommandCreator,
	commandRunner command_runner.CommandRunner,
	defaultRootfs string) *CompositeImagePlugin {

	return &CompositeImagePlugin{
		unprivilegedCommandCreator: unprivilegedCommandCreator,
		commandRunner:              commandRunner,
		defaultRootfs:              defaultRootfs,
	}
}

func (p *CompositeImagePlugin) Create(log lager.Logger, handle string, spec rootfs_provider.Spec) (string, []string, error) {
	log = log.Session("image-plugin-create")
	log.Debug("start")
	defer log.Debug("end")

	if spec.RootFS.String() == "" {
		var err error
		spec.RootFS, err = url.Parse(p.defaultRootfs)

		if err != nil {
			log.Error("parsing-default-rootfs-failed", err)
			return "", nil, errorwrapper.Wrap(err, "parsing default rootfs")
		}
	}

	createCmd, err := p.unprivilegedCommandCreator.CreateCommand(log, handle, spec)
	if err != nil {
		log.Error("create-image-plugin-command-failed", err)
		return "", nil, errorwrapper.Wrap(err, "creating image plugin create-command")
	}

	stdoutBuffer := bytes.NewBuffer([]byte{})
	createCmd.Stdout = stdoutBuffer
	createCmd.Stderr = lagregator.NewRelogger(log)

	if err := p.commandRunner.Run(createCmd); err != nil {
		logData := lager.Data{"action": "create", "stdout": stdoutBuffer.String()}
		log.Error("composite-image-plugin-result", err, logData)
		return "", nil, errorwrapper.Wrapf(err, "running image plugin create: %s", stdoutBuffer.String())
	}

	imagePath := strings.TrimSpace(stdoutBuffer.String())
	rootfsPath := filepath.Join(imagePath, "rootfs")

	envVars, err := p.readEnvVars(imagePath)
	if err != nil {
		log.Error("read-image-json-failed", err)
		return "", nil, errorwrapper.Wrap(err, "reading image.json")
	}

	return rootfsPath, envVars, nil
}

func (p *CompositeImagePlugin) Destroy(log lager.Logger, handle string, privileged bool) error {
	return nil
}

func (p *CompositeImagePlugin) Metrics(log lager.Logger, handle string, privileged bool) (garden.ContainerDiskStat, error) {
	return garden.ContainerDiskStat{}, nil
}

func (p *CompositeImagePlugin) GC(log lager.Logger) error {
	return nil
}

func (p *CompositeImagePlugin) readEnvVars(imagePath string) ([]string, error) {
	imageConfigFile, err := os.Open(filepath.Join(imagePath, "image.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}

		return nil, errorwrapper.Wrap(err, "could not open image configuration")
	}

	var imageConfig Image
	if err := json.NewDecoder(imageConfigFile).Decode(&imageConfig); err != nil {
		return nil, errorwrapper.Wrap(err, "parsing image config")
	}

	return imageConfig.Config.Env, nil
}
