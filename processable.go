package pluto

import (
	"time"
)

type Processable struct {
	Pipeline  string    `json:"pipeline"`
	Token     string    `json:"token,omitempty"`
	Body      any       `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
