package main_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	pathToMain   string
	pathToDocker string
	currentDir   string
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, ".")
}

var _ = BeforeSuite(func() {
	var (
		err              error
		valueIsSet       bool
		currentDirectory string
	)
	pathToMain, err = gexec.Build("github.com/predix/go-camel/examples/filetoconsole")
	Expect(err).NotTo(HaveOccurred())

	currentDirectory, valueIsSet = os.LookupEnv("PWD")
	Expect(valueIsSet).NotTo(BeFalse())
	currentDir, err = filepath.EvalSymlinks(currentDirectory)
	Expect(err).NotTo(HaveOccurred())

})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
