package pluto

import (
	"encoding/json"
	"net"
	"pluto/pkg/random"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	AcceptedConnections      = make(map[uuid.UUID]AcceptedConnection)
	AcceptedConnectionsMutex = new(sync.RWMutex)
)

const ConnectionTokenLength = 32

type AcceptedConnection struct {
	ID       uuid.UUID `json:"connection_id"`
	Token    string    `json:"connection_token"`
	net.Conn `json:"-"`
	Encoder  StreamEncoder `json:"-"`
}

var acceptor = NewInlineProcessor(func(processable Processable) (Processable, bool) {
	connection := AcceptedConnection{
		ID:      uuid.New(),
		Token:   random.String(ConnectionTokenLength),
		Conn:    processable.GetBody().(map[string]any)["connection"].(net.Conn),
		Encoder: NewChannelBasedStreamEncoder(NewJsonStreamEncoder(processable.GetBody().(map[string]any)["connection"].(net.Conn))),
	}

	b, err := json.Marshal(OutGoingProcessable{
		Consumer: ExternalIdentifier{
			Name: "CONNECTION_ACCEPTOR",
			Kind: KindPipeline,
		},
		Body: connection,
	})
	if err != nil {
		Log.Error("Marshal OutGoingProcessable", zap.Error(err))
		return processable, false
	}

	if err := connection.SetWriteDeadline(time.Now().Add(time.Second * 2)); err != nil {
		Log.Error("Set write deadline", zap.Error(err))
		return processable, false
	}

	if _, err := connection.Write(b); err != nil {
		Log.Debug("Write bytes to connection", zap.Error(err))
		return processable, false
	}

	AcceptedConnectionsMutex.Lock()
	AcceptedConnections[connection.ID] = connection
	AcceptedConnectionsMutex.Unlock()

	processable.GetBody().(map[string]any)["connection_id"] = connection.ID
	processable.GetBody().(map[string]any)["connection_token"] = connection.Token

	return processable, true
})
