package account

import (
	"net/http"
	"pluto"
)

type PasswordAuthenticator struct {
	Email    string
	Password string
}

func (p *PasswordAuthenticator) Authenticate() error {
	a, err := Find(p.Email)
	if err != nil {
		return &pluto.Error{
			HTTPCode: http.StatusUnauthorized,
			Message:  "The email or password is incorrect",
		}
	}

	if !a.Password.Compare([]byte(p.Password)) {
		return &pluto.Error{
			HTTPCode: http.StatusUnauthorized,
			Message:  "The email or password is incorrect",
		}
	}

	return nil
}
