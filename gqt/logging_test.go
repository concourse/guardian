package gqt_test

import (
	"code.cloudfoundry.org/garden"
	"github.com/concourse/guardian/gqt/runner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("runC Logging", func() {
	var (
		client   *runner.RunningGarden
		logLevel string
	)

	BeforeEach(func() {
		skipIfContainerd()
	})

	JustBeforeEach(func() {
		binPath, err := gexec.Build("code.cloudfoundry.org/guardian/gqt/cmd/fake_runc")
		Expect(err).NotTo(HaveOccurred())

		containerSpec := garden.ContainerSpec{}

		config.RuntimePluginBin = binPath
		config.LogLevel = logLevel
		client = runner.Start(config)
		client.Create(containerSpec)
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	Context("when the log-level is set to debug", func() {
		BeforeEach(func() {
			logLevel = "debug"
		})

		It("logs runc logs to stdout", func() {
			Eventually(client, "1s").Should(gbytes.Say("guardian-runc-logging-test-info"))
			Eventually(client, "1s").Should(gbytes.Say("guardian-runc-logging-test-warn"))
			Eventually(client, "1s").Should(gbytes.Say("guardian-runc-logging-test-error"))
			Eventually(client, "1s").Should(gbytes.Say("guardian-runc-logging-test-print"))
		})
	})

	Context("when the log-level is set to info", func() {
		BeforeEach(func() {
			logLevel = "info"
		})

		It("does not log runc logs to stdout", func() {
			Consistently(client, "1s").ShouldNot(gbytes.Say("guardian-runc-logging-test-info"))
			Consistently(client, "1s").ShouldNot(gbytes.Say("guardian-runc-logging-test-warn"))
			Consistently(client, "1s").ShouldNot(gbytes.Say("guardian-runc-logging-test-error"))
			Consistently(client, "1s").ShouldNot(gbytes.Say("guardian-runc-logging-test-print"))
		})
	})
})

var _ = Describe("garden server Logging", func() {
	var (
		client *runner.RunningGarden
	)

	JustBeforeEach(func() {
		client = runner.Start(config)
	})

	Context("when the log-level is passed as debug", func() {
		BeforeEach(func() {
			config.LogLevel = "debug"
		})

		AfterEach(func() {
			Expect(client.DestroyAndStop()).To(Succeed())
		})

		It("logs at debug level", func() {
			// create a container in order to execute a code path that prints debug
			// logs
			_, err := client.Create(garden.ContainerSpec{})
			Expect(err).NotTo(HaveOccurred())

			Eventually(client, "1s").Should(gbytes.Say(`"log_level":0`))
		})

		It("doesn't log spurious messages on start", func() {
			Consistently(client).ShouldNot(gbytes.Say(`guardian.iptables-runner.command.failed`))
			Consistently(client).ShouldNot(gbytes.Say(`guardian.failed-to-parse-port-pool-properties`))

			Eventually(client).Should(gbytes.Say(`guardian.no-port-pool-state-to-recover-starting-clean`))
		})
	})

	Context("when log-level is not passed", func() {
		BeforeEach(func() {
			config.LogLevel = ""
		})

		AfterEach(func() {
			Expect(client.Stop()).To(Succeed())
		})

		It("logs at info level", func() {
			Eventually(client, "1s").Should(gbytes.Say(`"log_level":1`))
		})

		It("does not log at debug level", func() {
			Consistently(client, "1s").ShouldNot(gbytes.Say(`"log_level":0`))
		})
	})

	Context("when a nonsense log-level is passed", func() {
		BeforeEach(func() {
			config.StartupExpectedToFail = true
			config.LogLevel = "nonsense-log-level"
		})

		It("exits with a code of 1", func() {
			Eventually(client, "1s").Should(gexec.Exit(1))
		})
	})
})
