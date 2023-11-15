package sdk

type Processable struct {
	Producer           Identifier `json:"producer"`
	Consumer           Identifier `json:"consumer"`
	ProducerCredential Credential `json:"producer_credential"`
	Body               any        `json:"body"`
}
