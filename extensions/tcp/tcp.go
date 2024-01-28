package tcp

import (
	"net"
	"pluto"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const MAXRequestPerConnection = 1000

var ConnectionHandler = pluto.Pipeline{
	"TCP_CONNECTION_HANDLER",
	pluto.ProcessorBucket{[]pluto.Processor{acceptor, authenticator, processor}},
}

func init() {
	go func() {
		l, err := net.Listen("tcp4", pluto.Env.TCPServerAddress)
		if err != nil {
			pluto.Log.Fatal("Create TCP listener", zap.String("address", pluto.Env.TCPServerAddress), zap.Error(err))
		}

		for {
			pluto.Log.Debug("Waiting for connections")

			// TODO: Any check or feature to accept new connections.

			conn, err := l.Accept()
			if err != nil {
				pluto.Log.Debug("Failed to accept new connection", zap.Error(err))
				continue
			}

			pluto.ApplicationLogger.Debug(pluto.ApplicationLog{
				Message: "New TCP connection accepted",
				Extra:   map[string]any{"remote_address": conn.RemoteAddr().String()},
			})

			go ConnectionHandler.Process(&pluto.InternalProcessable{
				ID:        uuid.New(),
				Body:      map[string]any{"connection": conn},
				CreatedAt: time.Now(),
			})
		}
	}()
}

var processor = pluto.NewFinalProcessor(
	&ConnectionDecoder{
		MaxDecode:    MAXRequestPerConnection,
		ReadDeadline: time.Hour,
		ProcessableBuilder: func(context pluto.Processable, new pluto.OutComingProcessable) pluto.Processable {
			defer func() { recover() }()

			AuthenticatedConnectionsMutex.RLock()
			defer AuthenticatedConnectionsMutex.RUnlock()

			connection := AuthenticatedConnections[context.GetBody().(map[string]any)["connection_id"].(uuid.UUID)]
			new.Producer = connection.Producer.(pluto.ExternalIdentifier)
			new.ProducerCredential = connection.ProducerCredential.(pluto.OutComingCredential)
			new.Connection = connection.AcceptedConnection
			new.Encoder = connection.AcceptedConnection.Encoder

			return &new
		},
		Processor: pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
			pluto.Process(processable.(pluto.RoutableProcessable))
			return processable, true
		}),
	},
).Final(
	pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
		AuthenticatedConnectionsMutex.Lock()
		defer AuthenticatedConnectionsMutex.Unlock()
		defer func() { recover() }()

		authenticatedConnection := AuthenticatedConnections[processable.GetBody().(map[string]any)["connection_id"].(uuid.UUID)]
		_ = authenticatedConnection.Encoder.Close()
		delete(AuthenticatedConnections, authenticatedConnection.ID)

		return processable, true
	}),
)
