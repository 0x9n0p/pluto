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
		// TODO:
		//  2. Risk of the DOS attack
		acceptor,

		&ConnectionDecoder{
			MaxDecode: 1,
			Processor: authenticator,
		},

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

			go ConnectionHandler.Process(&InternalProcessable{
				ID:        uuid.New(),
				Body:      Appendable{"connection": conn},
				CreatedAt: time.Now(),
			})
		}
	}()
}
