package pluto

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	AuthenticatedConnections      = make(map[uuid.UUID]AuthenticatedConnection)
	AuthenticatedConnectionsMutex = new(sync.RWMutex)
)

type AuthenticatedConnection struct {
	AcceptedConnection
	Producer           Identifier
	ProducerCredential Credential
}

var authenticator = NewConditionalProcessor(&ConnectionDecoder{
	MaxDecode:          1,
	ReadDeadline:       time.Second * 2,
	ProcessableBuilder: func(context Processable, new OutComingProcessable) Processable { return &new },
	Processor: NewInlineProcessor(func(processable Processable) (result Processable, succeed bool) {
		defer func() { succeed = recover() == nil }()

		ApplicationLogger.Debug(ApplicationLog{
			Message: "Authentication request",
		})

		var authenticatedConnection AuthenticatedConnection

		// Authenticate Connection
		{
			connectionID := uuid.MustParse(processable.GetBody().(map[string]any)["connection_id"].(string))
			connectionToken := processable.GetBody().(map[string]any)["connection_token"].(string)

			AcceptedConnectionsMutex.RLock()
			acceptedConnection, found := AcceptedConnections[connectionID]
			if !found || acceptedConnection.Token != connectionToken {
				AcceptedConnectionsMutex.RUnlock()
				return processable, false
			}
			AcceptedConnectionsMutex.RUnlock()

			authenticatedConnection.AcceptedConnection = acceptedConnection
		}

		// Authenticate Producer
		{
			p := processable.(*OutComingProcessable)

			validated, err := p.ProducerCredential.Validate(p.GetProducer())
			if err != nil {
				Log.Debug("Validate producer credential", zap.Error(err))
				return processable, false
			}

			if !validated {
				// TODO: Write to the client that the credential is not valid + Set write deadline
				return processable, false
			}

			authenticatedConnection.Producer = p.Producer
			authenticatedConnection.ProducerCredential = p.ProducerCredential
		}

		// Reset write deadline
		if err := authenticatedConnection.Conn.SetWriteDeadline(time.Time{}); err != nil {
			return processable, false
		}

		// Move the connection from the accepted connections to the authenticated connections
		{
			AcceptedConnectionsMutex.Lock()
			AuthenticatedConnectionsMutex.Lock()

			delete(AcceptedConnections, authenticatedConnection.AcceptedConnection.ID)
			AuthenticatedConnections[authenticatedConnection.ID] = authenticatedConnection

			AuthenticatedConnectionsMutex.Unlock()
			AcceptedConnectionsMutex.Unlock()
		}

		return processable, true
	}),
}).Fail(ProcessorBucket{Processors: []Processor{
	NewInlineProcessor(func(processable Processable) (Processable, bool) {
		defer func() { recover() }()

		connectionID := processable.GetBody().(map[string]any)["connection_id"].(uuid.UUID)
		connectionToken := processable.GetBody().(map[string]any)["connection_token"].(string)

		AcceptedConnectionsMutex.RLock()
		acceptedConnection, found := AcceptedConnections[connectionID]
		if !found || acceptedConnection.Token != connectionToken {
			AcceptedConnectionsMutex.RUnlock()
			return processable, false
		}
		AcceptedConnectionsMutex.RUnlock()

		_ = acceptedConnection.Close()

		AcceptedConnectionsMutex.Lock()
		delete(AcceptedConnections, acceptedConnection.ID)
		AcceptedConnectionsMutex.Unlock()

		return processable, false
	}),
}})
