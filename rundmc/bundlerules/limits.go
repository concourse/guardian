package bundlerules

import (
	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

var CpuPeriod uint64 = 100000
var MinCpuQuota uint64 = 1000

type Limits struct {
	CpuQuotaPerShare uint64
}

func (l Limits) Apply(bndl goci.Bndl, spec gardener.DesiredContainerSpec) goci.Bndl {
	limit := uint64(spec.Limits.Memory.LimitInBytes)
	bndl = bndl.WithMemoryLimit(specs.LinuxMemory{Limit: &limit, Swap: &limit})

	shares := uint64(spec.Limits.CPU.LimitInShares)
	cpuSpec := specs.LinuxCPU{Shares: &shares}
	if l.CpuQuotaPerShare > 0 && shares > 0 {
		cpuSpec.Period = &CpuPeriod

		quota := shares * l.CpuQuotaPerShare
		if quota < MinCpuQuota {
			quota = MinCpuQuota
		}
		cpuSpec.Quota = &quota
	}
	bndl = bndl.WithCPUShares(cpuSpec)

	pids := int64(spec.Limits.Pid.Max)
	return bndl.WithPidLimit(specs.LinuxPids{Limit: pids})
}
