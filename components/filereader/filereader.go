package filereader

import (
	"io/ioutil"

	"github.com/predix/go-camel/core"
)

type filereader struct {
	filepath string
}

func (f *filereader) Process(xchng core.Exchange) core.Message {
	fileContent, err := ioutil.ReadFile(f.filepath)
	if err != nil {
		return core.Message{
			Header: map[string]interface{}{
				"filereader-error": err.Error(),
			},
		}
	}
	return core.Message{
		Body: fileContent,
	}
}

func New(filepath string) (core.Endpoint, error) {
	return &filereader{
		filepath: filepath,
	}, nil
}
