package printer

import (
	"io"

	"github.com/predix/go-camel/core"
)

const bytesWritten = "printer-bytes-written"
const errorHeader = "printer-error"

type printer struct {
	writer io.Writer
}

func (p *printer) Process(xchng core.Exchange) core.Message {
	//TODO do the actual implementation
	n, err := p.writer.Write(xchng.In.Body)
	if err != nil {
		return core.Message{
			Header: map[string]interface{}{
				errorHeader: err.Error(),
			},
		}
	}

	return core.Message{
		Header: map[string]interface{}{
			bytesWritten: n,
		},
	}
}

func NewPrinter(w io.Writer) core.Endpoint {
	return &printer{
		writer: w,
	}
}
