package guardiancmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/commandrunner"
	"code.cloudfoundry.org/commandrunner/windows_command_runner"
	"github.com/concourse/guardian/gardener"
	"github.com/concourse/guardian/kawasaki"
	"github.com/concourse/guardian/rundmc"
	"github.com/concourse/guardian/rundmc/bundlerules"
	"github.com/concourse/guardian/rundmc/execrunner"
	"github.com/concourse/guardian/rundmc/goci"
	"github.com/concourse/guardian/rundmc/preparerootfs"
	"github.com/concourse/guardian/rundmc/processes"
	"github.com/concourse/guardian/rundmc/runcontainerd"
	"github.com/concourse/guardian/rundmc/runrunc"
	"github.com/concourse/guardian/rundmc/users"
	"code.cloudfoundry.org/lager"
	"github.com/docker/docker/pkg/mount"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type WindowsFactory struct {
	config        *ServerCommand
	commandRunner commandrunner.CommandRunner
}

func (cmd *ServerCommand) NewGardenFactory() GardenFactory {
	return &WindowsFactory{config: cmd, commandRunner: windows_command_runner.New(false)}
}

type NoopStarter struct{}

func (s *NoopStarter) Start() error {
	return nil
}

func (f *WindowsFactory) WireCgroupsStarter(_ lager.Logger) gardener.Starter {
	return &NoopStarter{}
}

func (cmd *SetupCommand) WireCgroupsStarter(_ lager.Logger) gardener.Starter {
	return &NoopStarter{}
}

type NoopResolvConfigurer struct{}

func (*NoopResolvConfigurer) Configure(log lager.Logger, cfg kawasaki.NetworkConfig, pid int) error {
	return nil
}

func (f *WindowsFactory) WireResolvConfigurer() kawasaki.DnsResolvConfigurer {
	return &NoopResolvConfigurer{}
}

func (f *WindowsFactory) WireVolumizer(logger lager.Logger) gardener.Volumizer {
	if f.config.Image.Plugin.Path() != "" || f.config.Image.PrivilegedPlugin.Path() != "" {
		return f.config.wireImagePlugin(f.commandRunner, 0, 0)
	}

	noop := gardener.NoopVolumizer{}
	return gardener.NewVolumeProvider(noop, noop, gardener.CommandFactory(preparerootfs.Command), f.commandRunner, 0, 0)
}

func (f *WindowsFactory) WireExecRunner(runMode, runcRoot string, _, _ uint32) runrunc.ExecRunner {
	return execrunner.NewWindowsExecRunner(f.config.Runtime.Plugin, runMode, f.commandRunner)
}

func (f *WindowsFactory) WireRootfsFileCreator() rundmc.RootfsFileCreator {
	return noopRootfsFileCreator{}
}

type noopRootfsFileCreator struct{}

func (noopRootfsFileCreator) CreateFiles(rootFSPath string, pathsToCreate ...string) error {
	return nil
}

func (f *WindowsFactory) CommandRunner() commandrunner.CommandRunner {
	return f.commandRunner
}

type mkdirer struct{}

func (m mkdirer) MkdirAs(rootFSPathFile string, uid, gid int, mode os.FileMode, recreate bool, paths ...string) error {
	for _, path := range paths {
		volumeName := filepath.VolumeName(path)
		if err := os.MkdirAll(filepath.Join(rootFSPathFile, strings.TrimPrefix(path, volumeName)), 0755); err != nil {
			return err
		}
	}
	return nil
}

func (f *WindowsFactory) WireMkdirer() runrunc.Mkdirer {
	return mkdirer{}
}

func wireEnvFunc() processes.EnvFunc {
	return processes.WindowsEnvFor
}

func wireMounts() bundlerules.Mounts {
	noopMountOptionsGetter := func(path string, mountInfos []*mount.Info) ([]string, error) {
		return []string{}, nil
	}
	noopMountInfosProvider := func() ([]*mount.Info, error) {
		return []*mount.Info{}, nil
	}
	return bundlerules.Mounts{MountOptionsGetter: noopMountOptionsGetter, MountInfosProvider: noopMountInfosProvider}
}

// Note - it's not possible to bind mount a single file in Windows, so we are
// using a directory instead
func initBindMountAndPath(initPathOnHost string) (specs.Mount, string) {
	initPathInContainer := filepath.Join(`C:\`, "Windows", "Temp", "bin", filepath.Base(initPathOnHost))
	return specs.Mount{
		Type:        "bind",
		Source:      filepath.Dir(initPathOnHost),
		Destination: filepath.Dir(initPathInContainer),
		Options:     []string{"bind"},
	}, initPathInContainer
}

func defaultBindMounts() []specs.Mount {
	return []specs.Mount{}
}

func privilegedMounts() []specs.Mount {
	return []specs.Mount{}
}

func unprivilegedMounts() []specs.Mount {
	return []specs.Mount{}
}

func getPrivilegedDevices() []specs.LinuxDevice {
	return nil
}

func bindMountPoints() []string {
	return nil
}

func mustGetMaxValidUID() int {
	return -1
}

func ensureServerSocketDoesNotLeak(socketFD uintptr) error {
	panic("this should be unreachable: no sockets on Windows")
}

func wireContainerd(socket string, bndlLoader *goci.BndlLoader, processBuilder *processes.ProcBuilder, userLookupper users.UserLookupper,
	wireExecer func(pidGetter runrunc.PidGetter) *runrunc.Execer, statser *runrunc.Statser, useContainerdForProcesses bool) (*runcontainerd.RunContainerd, *runcontainerd.RunContainerPea, *runcontainerd.PidGetter, error) {
	return nil, nil, nil, errors.New("containerd not impletemented on windows")
}

func containerdRuncRoot() string {
	return ""
}

func (cmd *ServerCommand) computeRuncRoot() string {
	return ""
}
