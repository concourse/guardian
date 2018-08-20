package bundlerules_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	specs "github.com/opencontainers/runtime-spec/specs-go"

	spec "github.com/concourse/guardian/gardener/container-spec"
	"github.com/concourse/guardian/rundmc/bundlerules"
	"github.com/concourse/guardian/rundmc/goci"
)

var _ = Describe("RootFS", func() {
	var (
		rule bundlerules.RootFS

		rootfsPath     string
		returnedBundle goci.Bndl
	)

	JustBeforeEach(func() {
		rootfsPath = "banana/"

		rule = bundlerules.RootFS{}

		var err error
		returnedBundle, err = rule.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			BaseConfig: specs.Spec{Root: &specs.Root{Path: rootfsPath}},
		}, "not-needed-path")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(rootfsPath)).To(Succeed())
	})

	It("applies the rootfs to the passed bundle", func() {
		Expect(returnedBundle.Spec.Root.Path).To(Equal(rootfsPath))
	})
})
