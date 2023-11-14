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
	Consumer ExternalIdentifier `json:"consumer"`
	Body     any                `json:"body"`
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

// InternalProcessable
// TODO: Do we need producer?
type InternalProcessable struct {
	ID        uuid.UUID  `json:"id"`
	Producer  Identifier `json:"producer"`
	Body      any        `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
}

func (o *InternalProcessable) GetProducer() Identifier {
	return o.Producer
}

func (o *InternalProcessable) SetBody(v any) {
	o.Body = v
}

func (o *InternalProcessable) GetBody() any {
	return o.Body
}

func (o *InternalProcessable) UniqueProperty() string {
	return o.ID.String()
}

func (o *InternalProcessable) PredefinedKind() string {
	return KindInternalProcessable
}
