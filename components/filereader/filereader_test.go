package filereader_test

import (
	"io/ioutil"
	"math/rand"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/predix/go-camel/components/filereader"
	"github.com/predix/go-camel/core"
)

var _ = Describe("FileReader", func() {

	Describe("Process", func() {
		var filename string
		var fileContent []byte

		BeforeEach(func() {
			tempFile, err := ioutil.TempFile("", "")
			Expect(err).NotTo(HaveOccurred())
			fileContent = []byte(strconv.Itoa(rand.Int()))
			err = ioutil.WriteFile(tempFile.Name(), fileContent, 644)
			Expect(err).NotTo(HaveOccurred())
			filename = tempFile.Name()
		})

		It("Reads the given file and outputs the message", func() {
			fileReader, err := filereader.New(filename)
			Expect(err).NotTo(HaveOccurred())

			msg := fileReader.Process(core.Exchange{})
			Expect(msg.Body).To(Equal(fileContent))
		})

		Context("failure cases", func() {
			It("adds the error header", func() {
				fileReader, err := filereader.New("/some/file/that/does/not/exist")
				Expect(err).NotTo(HaveOccurred())

				msg := fileReader.Process(core.Exchange{})
				Expect(msg.Header).To(HaveKeyWithValue("filereader-error", "open /some/file/that/does/not/exist: no such file or directory"))
			})
		})
	})
})
