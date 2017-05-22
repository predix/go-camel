package filereader_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFilereaderComponents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "filereadercomponent")
}
