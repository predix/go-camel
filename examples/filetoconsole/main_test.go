package main_test

import (
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("reads the sample file and pushes the content to console", func() {
		command := exec.Command(pathToMain,
			"-f", filepath.Join(currentDir, "fixtures", "sample.dat"))
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))

		Expect(session.Out.Contents()).To(ContainSubstring("Lorem ipsum dolor sit amet"))
	})

	Context("failure cases", func() {
		It("fails when executed with the file parameter is not given", func() {
			command := exec.Command(pathToMain)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("-f is a required argument"))
		})
	})
})
