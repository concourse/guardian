package factory

import (
	"os"

	"github.com/concourse/guardian/kawasaki"
	"github.com/concourse/guardian/kawasaki/configure"
	"github.com/concourse/guardian/kawasaki/devices"
	"github.com/concourse/guardian/kawasaki/dns"
	"github.com/concourse/guardian/kawasaki/iptables"
	"github.com/concourse/guardian/kawasaki/netns"
)

func NewDefaultConfigurer(ipt *iptables.IPTablesController, depotDir string) kawasaki.Configurer {
	resolvConfigurer := &kawasaki.ResolvConfigurer{
		HostsFileCompiler: &dns.HostsFileCompiler{},
		ResolvCompiler:    &dns.ResolvCompiler{},
		DepotDir:          depotDir,
		ResolvFilePath:    "/etc/resolv.conf",
	}

	hostConfigurer := &configure.Host{
		Veth:       &devices.VethCreator{},
		Link:       &devices.Link{},
		Bridge:     &devices.Bridge{},
		FileOpener: netns.Opener(os.Open),
	}

	containerConfigurer := &configure.Container{
		FileOpener: netns.Opener(os.Open),
	}

	return kawasaki.NewConfigurer(
		resolvConfigurer,
		hostConfigurer,
		containerConfigurer,
		iptables.NewInstanceChainCreator(ipt),
	)
}
