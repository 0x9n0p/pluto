package pluto

import (
	"sync"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	AuthenticatedConnections      = make(map[uuid.UUID]AuthenticatedConnection)
	AuthenticatedConnectionsMutex = new(sync.RWMutex)
)

type AuthenticatedConnection struct {
	AcceptedConnection
	ProducerCredential Credential
	// TODO: Authenticated connections can only be removed by expiration
}

var authenticator = NewConditionalProcessor(&ConnectionDecoder{
	MaxDecode: 1,
	Processor: NewInlineProcessor(func(processable Processable) (result Processable, succeed bool) {
		defer func() { succeed = recover() == nil }()

		ApplicationLogger.Debug(ApplicationLog{
			Message: "Authentication request",
		})

		var authenticatedConnection AuthenticatedConnection

		// Authenticate Connection
		{
			connectionID := uuid.MustParse(processable.GetBody().(Appendable)["connection_id"].(string))
			connectionToken := processable.GetBody().(Appendable)["connection_token"].(string)

			acceptedConnection, found := GetAcceptedConnection(connectionID)
			if !found || acceptedConnection.Token != connectionToken {
				return processable, false
			}

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
				// TODO: Write to the client that the credential is not valid
				return processable, false
			}

			authenticatedConnection.ProducerCredential = p.ProducerCredential
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
		connectionID := uuid.MustParse(processable.GetBody().(Appendable)["connection_id"].(string))
		connectionToken := processable.GetBody().(Appendable)["connection_token"].(string)

		acceptedConnection, found := GetAcceptedConnection(connectionID)
		if !found || acceptedConnection.Token != connectionToken {
			return processable, false
		}

		AcceptedConnectionsMutex.Lock()
		delete(AcceptedConnections, acceptedConnection.ID)
		AcceptedConnectionsMutex.Unlock()

		return processable, false
	}),
}})
