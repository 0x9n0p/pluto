package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/google/uuid"
)

var ErrInvalidCredential = errors.New("invalid credential")

type TCPClient struct {
	ServerAddress      string
	Producer           Identifier
	ProducerCredential Credential

	ConnectionID    uuid.UUID
	ConnectionToken string

	net.Conn
	*json.Decoder
}

func (c *TCPClient) Connect() (err error) {
	c.Conn, err = net.Dial("tcp4", c.ServerAddress)
	if err != nil {
		return fmt.Errorf("connect to tcp server: %v", err)
	}

	c.Decoder = json.NewDecoder(c.Conn)

	p, err := c.Receive()
	if err != nil {
		return err
	}

	defer func() {
		if recover() != nil {
			err = ErrInvalidCredential
		}
	}()

	c.ConnectionID = uuid.MustParse(p.Body.(map[string]any)["connection_id"].(string))
	c.ConnectionToken = p.Body.(map[string]any)["connection_token"].(string)

	if err := c.Send(Processable{
		Producer:           c.Producer,
		ProducerCredential: c.ProducerCredential,
		Body: map[string]any{
			"connection_id":    c.ConnectionID,
			"connection_token": c.ConnectionToken,
		},
	}); err != nil {
		return fmt.Errorf("send credentials: %v", err)
	}

	return
}

func (c *TCPClient) Send(processable Processable) error {
	b, err := json.Marshal(processable)
	if err != nil {
		return fmt.Errorf("marshal processable: %v", err)
	}

	if _, err := c.Write(b); err != nil {
		return fmt.Errorf("write to connection: %v", err)
	}

	return nil
}

func (c *TCPClient) Receive() (p Processable, err error) {
	if err := c.Decode(&p); err != nil {
		return Processable{}, fmt.Errorf("decoder: %v", err)
	}
	return
}
