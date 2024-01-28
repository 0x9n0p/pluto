package tcp

import (
	"encoding/json"
	"io"
	"pluto"
)

type StreamEncoder interface {
	Encode(any) error
	Close() error
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

func (e *JsonStreamEncoder) Close() error {
	return nil
}

var DefaultChannelBasedStreamEncoder_OnError = func(err error) {
	pluto.ApplicationLogger.Debug(pluto.ApplicationLog{
		Message: "Stream encoder failed",
		Extra: map[string]any{
			"issuer": "ChannelBasedStreamEncoder",
			"error":  err.Error(),
		},
	})
}

type ChannelBasedStreamEncoder struct {
	channel chan any
	OnError func(error)
}

func NewChannelBasedStreamEncoder(encoder StreamEncoder) *ChannelBasedStreamEncoder {
	cencoder := &ChannelBasedStreamEncoder{
		channel: make(chan any),
		OnError: DefaultChannelBasedStreamEncoder_OnError,
	}

	go func() {
		for {
			select {
			case r, ok := <-cencoder.channel:
				if !ok {
					return
				}

				if err := encoder.Encode(r); err != nil {
					cencoder.OnError(err)
				}
			}
		}
	}()

	return cencoder
}

func (c *ChannelBasedStreamEncoder) Encode(a any) error {
	c.channel <- a
	return nil
}

func (c *ChannelBasedStreamEncoder) Close() error {
	close(c.channel)
	return nil
}
