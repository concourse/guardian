// +build !linux

package factory

import (
	"github.com/concourse/guardian/kawasaki"
	"github.com/concourse/guardian/kawasaki/iptables"
)

func NewDefaultConfigurer(ipt *iptables.IPTablesController, depotDir string) kawasaki.Configurer {
	panic("not supported on this platform")
}
