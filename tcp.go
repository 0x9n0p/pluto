package pluto

import (
	"net"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Address
// TODO: env
const Address = "localhost:9631"
const MAXRequestPerConnection = 1000

var Listener = func() net.Listener {
	l, err := net.Listen("tcp4", Address)
	if err != nil {
		Log.Fatal("Create TCP listener", zap.String("address", Address))
	}
	return l
}()

var ConnectionHandler = Pipeline{
	Name: "TCP_CONNECTION_HANDLER",
	ProcessorBucket: ProcessorBucket{Processors: []Processor{
		acceptor,
		authenticator,

		NewFinalProcessor(
			&ConnectionDecoder{
				MaxDecode:    MAXRequestPerConnection,
				ReadDeadline: time.Hour,
				Processor: NewInlineProcessor(func(processable Processable) (Processable, bool) {
					Process(processable.(RoutableProcessable))
					return processable, true
				}),
			},
		).Final(
			NewInlineProcessor(func(processable Processable) (Processable, bool) {
				AuthenticatedConnectionsMutex.Lock()
				defer AuthenticatedConnectionsMutex.Unlock()
				defer func() { recover() }()
				delete(AuthenticatedConnections, processable.GetBody().(map[string]any)["connection_id"].(uuid.UUID))
				return processable, true
			}),
		),
	}},
}

func init() {
	go func() {
		for {
			Log.Debug("Waiting for connections")

			// TODO: Any check or feature to accept new connections.

			conn, err := Listener.Accept()
			if err != nil {
				Log.Debug("Failed to accept new connection", zap.Error(err))
				continue
			}

			go ConnectionHandler.Process(&InternalProcessable{
				ID:        uuid.New(),
				Body:      map[string]any{"connection": conn},
				CreatedAt: time.Now(),
			})
		}
	}()
}
