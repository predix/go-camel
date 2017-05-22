package printer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrinterComponents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "printercomponent")
}
