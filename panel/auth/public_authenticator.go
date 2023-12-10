package auth

import (
	"net/http"
	"pluto"
)

type PublicKeyAuthenticator struct {
	SignedMessage string `json:"signed_message"`
}

// Authenticate
// TODO
func (p *PublicKeyAuthenticator) Authenticate() error {
	return &pluto.Error{
		HTTPCode: http.StatusNotImplemented,
		Message:  "Public key authenticator not implemented",
	}
}
