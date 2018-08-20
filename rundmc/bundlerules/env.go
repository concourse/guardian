package bundlerules

import (
	spec "github.com/concourse/guardian/gardener/container-spec"
	"github.com/concourse/guardian/rundmc/goci"
)

type Env struct {
}

func (r Env) Apply(bndl goci.Bndl, spec spec.DesiredContainerSpec, _ string) (goci.Bndl, error) {
	process := bndl.Process()
	var baseEnv []string
	if spec.BaseConfig.Process != nil {
		baseEnv = spec.BaseConfig.Process.Env
	}
	process.Env = append(baseEnv, spec.Env...)
	return bndl.WithProcess(process), nil
}
