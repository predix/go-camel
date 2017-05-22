package printer_test

import (
	"errors"

	"github.com/predix/go-camel/components/printer"
	"github.com/predix/go-camel/core"
	"github.com/predix/go-camel/fake"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrinterComponent", func() {
	var fakeWriter *fake.Writer
	var xchng core.Exchange

	BeforeEach(func() {
		fakeWriter = fake.NewFakeWriter()
		xchng = core.Exchange{
			In: core.Message{
				Body: []byte("SOME_BODY"),
			},
		}
	})

	Describe("Process", func() {
		It("prints the output to the output we choose to", func() {

			fakeWriter.WriteCall.Returns.Count = 123

			consolePrinter := printer.NewPrinter(fakeWriter)

			msg := consolePrinter.Process(xchng)
			Expect(fakeWriter.WriteCall.CallCount).To(Equal(1))
			Expect(fakeWriter.WriteCall.Recieves.Bytes).To(Equal([]byte("SOME_BODY")))

			Expect(msg.Header).To(HaveKeyWithValue("printer-bytes-written", 123))
		})

		Context("failure cases", func() {
			It("adds the error info to the header when the writing fails", func() {
				fakeWriter.WriteCall.Returns.Error = errors.New("some_error_occured")
				consolePrinter := printer.NewPrinter(fakeWriter)

				msg := consolePrinter.Process(xchng)
				Expect(fakeWriter.WriteCall.CallCount).To(Equal(1))
				Expect(fakeWriter.WriteCall.Recieves.Bytes).To(Equal([]byte("SOME_BODY")))
				//
				Expect(msg.Header).To(HaveKeyWithValue("printer-error", "some_error_occured"))
			})
		})
	})
})
