package fake

type Writer struct {
	WriteCall struct {
		CallCount int
		Recieves  struct {
			Bytes []byte
		}
		Returns struct {
			Error error
			Count int
		}
	}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.WriteCall.CallCount++
	w.WriteCall.Recieves.Bytes = p
	return w.WriteCall.Returns.Count, w.WriteCall.Returns.Error
}

func NewFakeWriter() *Writer {
	return &Writer{}
}
