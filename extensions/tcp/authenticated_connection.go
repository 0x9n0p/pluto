package tcp

import (
	"pluto"
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
	Producer           pluto.Identifier
	ProducerCredential pluto.Credential
}

var authenticator = pluto.NewConditionalProcessor(&ConnectionDecoder{
	MaxDecode:          1,
	ReadDeadline:       time.Second * 2,
	ProcessableBuilder: func(context pluto.Processable, new pluto.OutComingProcessable) pluto.Processable { return &new },
	Processor: pluto.NewInlineProcessor(func(processable pluto.Processable) (result pluto.Processable, succeed bool) {
		defer func() { succeed = recover() == nil }()

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
			p := processable.(*pluto.OutComingProcessable)

			validated, err := p.ProducerCredential.Validate(p.GetProducer())
			if err != nil {
				pluto.Log.Debug("Validate producer credential", zap.Error(err))
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

		pluto.ApplicationLogger.Debug(pluto.ApplicationLog{
			Message: "New connection successfully authenticated",
			Extra: map[string]any{
				"connection_id":  authenticatedConnection.ID,
				"remote_address": authenticatedConnection.RemoteAddr().String(),
			},
		})

		return processable, true
	}),
}).Fail(pluto.ProcessorBucket{Processors: []pluto.Processor{
	pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
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
		_ = acceptedConnection.Encoder.Close()

		AcceptedConnectionsMutex.Lock()
		delete(AcceptedConnections, acceptedConnection.ID)
		AcceptedConnectionsMutex.Unlock()

		return processable, false
	}),
}})
