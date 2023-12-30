package account

import (
	"net/http"
	"pluto"
)

type PasswordAuthenticator struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *PasswordAuthenticator) Authenticate() error {
	if p.Email != "admin" || p.Password != "admin" {
		return &pluto.Error{
			HTTPCode: http.StatusUnauthorized,
			Message:  "The email or password is incorrect",
		}
	}
	return nil
}
