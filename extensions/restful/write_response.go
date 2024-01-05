package restful

import (
	"pluto"
)

const ProcessorName_WriteResponse = "WRITE_RESPONSE"

func creator_WriteResponse(args []pluto.Value) (p pluto.Processor, err error) {
	return nil, nil
}

type WriteResponse struct {
	HTTPCode int
	Body     any
}

func (w *WriteResponse) Process(processable pluto.Processable) (pluto.Processable, bool) {
	// TODO: Get the HTTP connection and write the HTTPCode and the body if was not empty.
	return processable, true
}
