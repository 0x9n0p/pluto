package pluto

import (
	"time"

	"github.com/google/uuid"
)

type Processable interface {
	GetProducer() Identifier
	GetBody() any
}

type RoutableProcessable interface {
	Processable
	GetConsumer() Identifier
}

type OutComingProcessable struct {
	Producer           Identifier
	Consumer           Identifier
	ProducerCredential Credential
	Body               any
}

func (o *OutComingProcessable) GetProducer() Identifier {
	return o.Producer
}

func (o *OutComingProcessable) GetConsumer() Identifier {
	return o.Consumer
}

func (o *OutComingProcessable) GetBody() any {
	return o.Body
}

type OutGoingProcessable struct {
	Producer           Identifier
	Consumer           Identifier
	ProducerCredential Credential
	Body               any
}

func (o *OutGoingProcessable) GetProducer() Identifier {
	return o.Producer
}

func (o *OutGoingProcessable) GetConsumer() Identifier {
	return o.Consumer
}

func (o *OutGoingProcessable) GetBody() any {
	return o.Body
}

type InternalProcessable struct {
	ID        uuid.UUID  `json:"id"`
	Producer  Identifier `json:"producer"`
	Body      any        `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
}

func (i *InternalProcessable) GetProducer() Identifier {
	return i.Producer
}

func (i *InternalProcessable) GetBody() any {
	return i.Body
}

func (i *InternalProcessable) UniqueProperty() string {
	return i.ID.String()
}

func (i *InternalProcessable) PredefinedKind() string {
	return KindInternalProcessable
}
