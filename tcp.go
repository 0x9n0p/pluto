package pluto

import (
	"net"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const MAXRequestPerConnection = 1000

var ConnectionHandler = Pipeline{
	"TCP_CONNECTION_HANDLER",
	ProcessorBucket{[]Processor{acceptor, authenticator, processor}},
}

func init() {
	go func() {
		l, err := net.Listen("tcp4", Env.TCPServerAddress)
		if err != nil {
			Log.Fatal("Create TCP listener", zap.String("address", Env.TCPServerAddress))
		}

		for {
			Log.Debug("Waiting for connections")

			// TODO: Any check or feature to accept new connections.

			conn, err := l.Accept()
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

var processor = NewFinalProcessor(
	&ConnectionDecoder{
		MaxDecode:    MAXRequestPerConnection,
		ReadDeadline: time.Hour,
		ProcessableBuilder: func(context Processable, new OutComingProcessable) Processable {
			defer func() { recover() }()

			AuthenticatedConnectionsMutex.RLock()
			defer AuthenticatedConnectionsMutex.RUnlock()

			connection := AuthenticatedConnections[context.GetBody().(map[string]any)["connection_id"].(uuid.UUID)]
			new.Producer = connection.Producer.(ExternalIdentifier)
			new.ProducerCredential = connection.ProducerCredential.(OutComingCredential)
			new.Connection = connection.AcceptedConnection
			new.Encoder = connection.AcceptedConnection.Encoder

			return &new
		},
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
)
