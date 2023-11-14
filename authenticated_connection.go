package pluto

import (
	"sync"

	"github.com/google/uuid"
)

var (
	AuthenticatedConnections      = make([]AuthenticatedConnection, 0)
	AuthenticatedConnectionsMutex = new(sync.RWMutex)
)

type AuthenticatedConnection struct {
	AcceptedConnection
	Credential
	// TODO: expires
}

func GetAuthenticatedConnection(id uuid.UUID) (AuthenticatedConnection, bool) {
	for _, connection := range AuthenticatedConnections {
		if connection.ID == id {
			return connection, true
		}
	}
	return AuthenticatedConnection{}, false
}
