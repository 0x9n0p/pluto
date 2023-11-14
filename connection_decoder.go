package pluto

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConnectionDecoder struct {
	MaxDecode uint64
	Processor Processor
}

func (p ConnectionDecoder) Process(processable Processable) (Processable, bool) {
	decoder := json.NewDecoder(processable.GetBody().(Appendable)["connection"].(io.Reader))
	decoder.UseNumber()

	for i := uint64(0); i < p.MaxDecode; i++ {
		var outComingProcessable OutComingProcessable
		if err := decoder.Decode(&outComingProcessable); err != nil {
			Log.Debug("Decoding out-coming processable", zap.Error(err))
			return processable, false
		}

		if result, success := p.Processor.Process(&outComingProcessable); !success {
			return result, false
		}
	}

	return &InternalProcessable{
		ID:        uuid.New(),
		Body:      processable.GetBody(),
		CreatedAt: time.Now(),
	}, true
}

func (p ConnectionDecoder) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "CONNECTION_DECODER_PROCESSOR",
		Description: "",
		Input:       "",
		Output:      "",
	}
}
