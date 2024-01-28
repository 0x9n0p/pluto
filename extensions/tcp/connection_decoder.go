package tcp

import (
	"encoding/json"
	"net"
	"pluto"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConnectionDecoder struct {
	MaxDecode          uint64
	ReadDeadline       time.Duration
	ProcessableBuilder func(context pluto.Processable, new pluto.OutComingProcessable) pluto.Processable
	Processor          pluto.Processor
}

func (p ConnectionDecoder) Process(processable pluto.Processable) (pluto.Processable, bool) {
	conn := processable.GetBody().(map[string]any)["connection"].(net.Conn)

	decoder := json.NewDecoder(conn)
	decoder.UseNumber()

	for i := uint64(0); i < p.MaxDecode; i++ {
		if err := conn.SetReadDeadline(time.Now().Add(p.ReadDeadline)); err != nil {
			pluto.Log.Error("Set read deadline", zap.Error(err))
			return processable, false
		}

		var outComingProcessable pluto.OutComingProcessable
		if err := decoder.Decode(&outComingProcessable); err != nil {
			pluto.Log.Debug("Decoding out-coming processable", zap.Error(err))
			return processable, false
		}

		if result, success := p.Processor.Process(p.ProcessableBuilder(processable, outComingProcessable)); !success {
			return result, false
		}
	}

	return &pluto.InternalProcessable{
		ID:        uuid.New(),
		Body:      processable.GetBody(),
		CreatedAt: time.Now(),
	}, true
}
