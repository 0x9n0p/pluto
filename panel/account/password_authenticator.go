package account

import (
	"net/http"
	"pluto"
	"pluto/panel/database"
)

type PasswordAuthenticator struct {
	Email    string
	Password string
}

func (p *PasswordAuthenticator) Authenticate() (err error) {
	tx, err := database.Get().NewTransaction(false)
	if err != nil {
		return err
	}

	defer func() {
		if err2 := tx.Rollback(); err2 != nil {
			err = err2
		}
	}()

	a, err := Find(tx, p.Email)
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
