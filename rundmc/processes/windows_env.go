package processes

import (
	"code.cloudfoundry.org/garden"
	"github.com/concourse/guardian/rundmc/goci"
)

func WindowsEnvFor(bndl goci.Bndl, spec garden.ProcessSpec, _ int) []string {
	return append(bndl.Spec.Process.Env, spec.Env...)
}
