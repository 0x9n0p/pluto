package pluto

import (
	"encoding/json"
	"net"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConnectionDecoder struct {
	MaxDecode          uint64
	ReadDeadline       time.Duration
	ProcessableBuilder func(context Processable, new OutComingProcessable) Processable
	Processor          Processor
}

func (p ConnectionDecoder) Process(processable Processable) (Processable, bool) {
	conn := processable.GetBody().(map[string]any)["connection"].(net.Conn)

	decoder := json.NewDecoder(conn)
	decoder.UseNumber()

	for i := uint64(0); i < p.MaxDecode; i++ {
		if err := conn.SetReadDeadline(time.Now().Add(p.ReadDeadline)); err != nil {
			Log.Error("Set read deadline", zap.Error(err))
			return processable, false
		}

		var outComingProcessable OutComingProcessable
		if err := decoder.Decode(&outComingProcessable); err != nil {
			Log.Debug("Decoding out-coming processable", zap.Error(err))
			return processable, false
		}

		if result, success := p.Processor.Process(p.ProcessableBuilder(processable, outComingProcessable)); !success {
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
