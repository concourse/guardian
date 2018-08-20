package runcontainerd_test

import (
	"errors"
	"strings"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/guardian/rundmc/runcontainerd"
	"github.com/concourse/guardian/rundmc/runcontainerd/runcontainerdfakes"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RunContainerPea", func() {

	var (
		fakePeaManager     *runcontainerdfakes.FakePeaManager
		fakeProcessManager *runcontainerdfakes.FakeProcessManager

		runContainerPea *runcontainerd.RunContainerPea
	)

	BeforeEach(func() {
		fakePeaManager = new(runcontainerdfakes.FakePeaManager)
		fakeProcessManager = new(runcontainerdfakes.FakeProcessManager)

		runContainerPea = &runcontainerd.RunContainerPea{
			PeaManager:     fakePeaManager,
			ProcessManager: fakeProcessManager,
		}
	})

	Describe("Run", func() {
		It("creates a container using process args", func() {
			runContainerPea.Run(lagertest.NewTestLogger("test-logger"), "test-process-id", "/process/path", "", "", garden.ProcessIO{}, false, strings.NewReader(""), nil)

			Expect(fakePeaManager.CreateCallCount()).To(Equal(1))
			_, actualProcessPath, actualProcessID, _ := fakePeaManager.CreateArgsForCall(0)
			Expect(actualProcessPath).To(Equal("/process/path"))
			Expect(actualProcessID).To(Equal("test-process-id"))
		})

		When("the creator returns an error", func() {
			BeforeEach(func() {
				fakePeaManager.CreateReturns(errors.New("error-creating"))
			})

			It("returns the error", func() {
				_, err := runContainerPea.Run(lagertest.NewTestLogger("test-logger"), "test-process-id", "/process/path", "", "", garden.ProcessIO{}, false, strings.NewReader(""), nil)

				Expect(err).To(MatchError("error-creating"))
			})
		})
	})
})
