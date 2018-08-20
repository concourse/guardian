package bundlerules

import (
	spec "github.com/concourse/guardian/gardener/container-spec"
	"github.com/concourse/guardian/rundmc/goci"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type Namespaces struct{}

func (n Namespaces) Apply(bndl goci.Bndl, spec spec.DesiredContainerSpec, containerDir string) (goci.Bndl, error) {
	for ns, path := range spec.Namespaces {
		bndl = bndl.WithNamespace(specs.LinuxNamespace{Type: specs.LinuxNamespaceType(ns), Path: path})
	}

	return bndl, nil
}
