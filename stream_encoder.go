package pluto

import (
	"encoding/json"
	"io"
)

type StreamEncoder interface {
	Encode(any) error
}

type JsonStreamEncoder struct {
	*json.Encoder
}

func NewJsonStreamEncoder(writer io.Writer) *JsonStreamEncoder {
	return &JsonStreamEncoder{
		json.NewEncoder(writer),
	}
}

func (e *JsonStreamEncoder) Encode(a any) error {
	// TODO: Any application log
	return e.Encoder.Encode(a)
}
