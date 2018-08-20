package bundlerules_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/garden"
	spec "github.com/concourse/guardian/gardener/container-spec"
	"github.com/concourse/guardian/rundmc/bundlerules"
	"github.com/concourse/guardian/rundmc/goci"
)

var _ = Describe("LimitsRule", func() {
	It("sets the correct memory limit in bundle resources", func() {
		newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			Limits: garden.Limits{
				Memory: garden.MemoryLimits{LimitInBytes: 4096},
			},
		}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(*(newBndl.Resources().Memory.Limit)).To(BeNumerically("==", 4096))
	})

	It("limits swap to regular memory limit in bundle resources", func() {
		newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			Limits: garden.Limits{
				Memory: garden.MemoryLimits{LimitInBytes: 4096},
			},
		}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(newBndl.Resources().Memory.Swap).ToNot(BeNil())
		Expect(*(newBndl.Resources().Memory.Swap)).To(BeNumerically("==", 4096))
	})

	Context("when swap limit is disabled", func() {
		It("does not limit swap in bundle resources", func() {
			limits := bundlerules.Limits{DisableSwapLimit: true}
			newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					Memory: garden.MemoryLimits{LimitInBytes: 4096},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(newBndl.Resources().Memory.Swap).To(BeNil())
		})
	})

	It("sets the correct TCPMemoryLimit in the bundle resources", func() {
		limits := bundlerules.Limits{
			TCPMemoryLimit: 100,
		}
		newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(*(newBndl.Resources().Memory.KernelTCP)).To(Equal(limits.TCPMemoryLimit))
	})

	It("sets the provided BlockIOWeight in the bundle resources", func() {
		limits := bundlerules.Limits{
			BlockIOWeight: 100,
		}
		newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(*(newBndl.Resources().BlockIO.Weight)).To(Equal(limits.BlockIOWeight))
	})

	It("sets the correct CPU limit in bundle resources", func() {
		newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			Limits: garden.Limits{
				CPU: garden.CPULimits{Weight: 1},
			},
		}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(*(newBndl.Resources().CPU.Shares)).To(BeNumerically("==", 1))
		Expect(newBndl.Resources().CPU.Period).To(BeNil())
		Expect(newBndl.Resources().CPU.Quota).To(BeNil())
	})

	Context("when a positive cpu quota period per share is provided", func() {
		It("sets the correct CPU limit in bundle resources", func() {
			var quotaPerShare, weight uint64 = 100, 128
			limits := bundlerules.Limits{
				CpuQuotaPerShare: quotaPerShare,
			}
			newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					CPU: garden.CPULimits{Weight: weight},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Period)).To(BeNumerically("==", 100000))
			Expect(*(newBndl.Resources().CPU.Quota)).To(BeNumerically("==", weight*quotaPerShare))
		})
	})

	Context("when cpu quota * period per share is less than min valid cpu quota", func() {
		It("sets the min valid value of cpu quota in bundle resources", func() {
			limits := bundlerules.Limits{
				CpuQuotaPerShare: 1,
			}
			newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					CPU: garden.CPULimits{Weight: 1},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Quota)).To(BeNumerically("==", 1000))
		})
	})

	Context("when a zero cpu quota period per share is provided", func() {
		It("sets the correct CPU limit in bundle resources", func() {
			limits := bundlerules.Limits{
				CpuQuotaPerShare: 0,
			}
			newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					CPU: garden.CPULimits{Weight: 1},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Shares)).To(BeNumerically("==", 1))
			Expect(newBndl.Resources().CPU.Period).To(BeNil())
			Expect(newBndl.Resources().CPU.Quota).To(BeNil())
		})
	})

	Context("with positive cpu quota period per share and no shares", func() {
		It("sets the correct CPU limit in bundle resources", func() {
			limits := bundlerules.Limits{
				CpuQuotaPerShare: 5,
			}
			newBndl, err := limits.Apply(goci.Bundle(), spec.DesiredContainerSpec{}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Shares)).To(BeNumerically("==", 0))
			Expect(newBndl.Resources().CPU.Period).To(BeNil())
			Expect(newBndl.Resources().CPU.Quota).To(BeNil())
		})
	})

	Context("when LimitInShares is set", func() {
		It("sets the CPU shares", func() {
			newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					CPU: garden.CPULimits{LimitInShares: 1},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Shares)).To(BeNumerically("==", 1))
		})
	})

	Context("when both Weight and LimitInShares are set", func() {
		It("Weight has precedence ", func() {
			newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
				Limits: garden.Limits{
					CPU: garden.CPULimits{LimitInShares: 1, Weight: 2},
				},
			}, "not-needed-path")
			Expect(err).NotTo(HaveOccurred())

			Expect(*(newBndl.Resources().CPU.Shares)).To(BeNumerically("==", 2))
		})
	})

	It("sets the correct PID limit in bundle resources", func() {
		newBndl, err := bundlerules.Limits{}.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			Limits: garden.Limits{
				Pid: garden.PidLimits{Max: 1},
			},
		}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())

		Expect(newBndl.Resources().Pids.Limit).To(BeNumerically("==", 1))
	})
})
