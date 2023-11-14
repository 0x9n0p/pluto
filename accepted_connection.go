package pluto

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

var (
	AcceptedConnections      = make([]AcceptedConnection, 0)
	AcceptedConnectionsMutex = new(sync.RWMutex)
)

type AcceptedConnection struct {
	ID uuid.UUID
	net.Conn
	// TODO: expires
}

func GetAcceptedConnection(id uuid.UUID) (AcceptedConnection, bool) {
	for _, connection := range AcceptedConnections {
		if connection.ID == id {
			return connection, true
		}
	}
	return AcceptedConnection{}, false
}
