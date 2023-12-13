package pluto

import (
	"time"

	"github.com/google/uuid"
)

type Processable interface {
	SetBody(any)
	GetBody() any
}

type RoutableProcessable interface {
	Processable
	GetProducer() Identifier
	GetConsumer() Identifier
}

type OutComingProcessable struct {
	Producer           ExternalIdentifier  `json:"producer"`
	Consumer           ExternalIdentifier  `json:"consumer"`
	ProducerCredential OutComingCredential `json:"producer_credential"`
	Connection         AcceptedConnection  `json:"connection"`
	Body               any                 `json:"body"`
}

func (o *OutComingProcessable) GetProducer() Identifier {
	return o.Producer
}

func (o *OutComingProcessable) GetConsumer() Identifier {
	return o.Consumer
}

func (o *OutComingProcessable) SetBody(v any) {
	o.Body = v
}

func (o *OutComingProcessable) GetBody() any {
	return o.Body
}

type OutGoingProcessable struct {
	Consumer Identifier `json:"consumer"`
	Body     any        `json:"body"`
}

func (o *OutGoingProcessable) GetConsumer() Identifier {
	return o.Consumer
}

func (o *OutGoingProcessable) SetBody(v any) {
	o.Body = v
}

func (o *OutGoingProcessable) GetBody() any {
	return o.Body
}

type InternalProcessable struct {
	ID        uuid.UUID `json:"id"`
	Body      any       `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (o *InternalProcessable) SetBody(v any) {
	o.Body = v
}

func (o *InternalProcessable) GetBody() any {
	return o.Body
}
