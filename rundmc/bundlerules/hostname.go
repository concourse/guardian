package bundlerules

import (
	spec "github.com/concourse/guardian/gardener/container-spec"
	"github.com/concourse/guardian/rundmc/goci"
)

type Hostname struct {
}

func (r Hostname) Apply(bndl goci.Bndl, spec spec.DesiredContainerSpec, _ string) (goci.Bndl, error) {
	hostname := spec.Hostname
	if len(hostname) > 49 {
		hostname = hostname[len(hostname)-49:]
	}

	return bndl.WithHostname(hostname), nil
}
