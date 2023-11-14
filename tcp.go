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

// ConnectionHandler
// decoder -> authenticator -> loop{decoder -> execution cache} -> close_connection
// TODO: This pipeline can be configured using HTTP APIs
var ConnectionHandler = Pipeline{
	Name: "TCP_CONNECTION_HANDLER",
	ProcessorBucket: ProcessorBucket{Processors: []Processor{
		// TODO: Close and Remove connection when the processor returns false

		// TODO: Remove the connection from accepted connections
		&ConnectionDecoder{
			MaxDecode: 1,
			// TODO: Authenticator
			Processor: NewInlineProcessor(func(processable Processable) (Processable, bool) {
				ApplicationLogger.Debug(ApplicationLog{
					Message: "Authentication request",
					Extra:   map[string]any{},
				})
				return processable, true
			}),
		},

		// TODO: Move conn from accepted connections to authenticated connections
		NewInlineProcessor(func(processable Processable) (Processable, bool) {
			return processable, true
		}),

		// TODO: Remove connection from authenticated connections
		&ConnectionDecoder{
			MaxDecode: MAXRequestPerConnection,
			Processor: NewInlineProcessor(func(processable Processable) (Processable, bool) {
				Process(processable.(RoutableProcessable))
				return processable, true
			}),
		},
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

			acceptedConnection := AcceptedConnection{ID: uuid.New(), Conn: conn}

			AcceptedConnectionsMutex.Lock()
			AcceptedConnections = append(AcceptedConnections, acceptedConnection)
			AcceptedConnectionsMutex.Unlock()

			Log.Debug("New connection accepted", zap.String("remote_address", conn.RemoteAddr().String()))
			ApplicationLogger.Debug(ApplicationLog{
				Message: "New connection accepted",
				Extra:   map[string]any{"remote_address": conn.RemoteAddr().String()},
			})

			go ConnectionHandler.Process(&InternalProcessable{
				ID: uuid.New(),
				//Producer:  nil,
				Body:      Appendable{"connection": acceptedConnection},
				CreatedAt: time.Now(),
			})
		}
	}()
}
