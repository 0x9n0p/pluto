package account

import (
	"encoding/json"
	"net/http"
	"pluto"
	"pluto/panel/database"
	"time"

	"go.uber.org/zap"
)

type Account struct {
	Email    string    `json:"email"`
	Password Password  `json:"password"`
	SavedAt  time.Time `json:"saved_at"`

	Transaction *database.Transaction `json:"-"`
}

// ChangeEmail
// TODO
func (a *Account) ChangeEmail(new string) error {
	return &pluto.Error{
		HTTPCode: http.StatusNotImplemented,
		Message:  "Not Implemented",
	}
}

func (a *Account) ChangePassword(old, new []byte) (err error) {
	if !a.Password.Compare(old) {
		return &pluto.Error{
			HTTPCode: http.StatusUnauthorized,
			Message:  "The entered old password is incorrect",
		}
	}

	a.Password, err = NewPassword(new)
	if err != nil {
		return &pluto.Error{
			HTTPCode: http.StatusBadRequest,
			Message:  "Make sure you have entered the new password in the correct format",
		}
	}

	return a.Save()
}

func (a *Account) Save() error {
	b, err := json.Marshal(a)
	if err != nil {
		pluto.Log.Error("Failed to marshal the account", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	err = a.Transaction.Bucket(bucket).Put([]byte(a.Email), b)
	if err != nil {
		pluto.Log.Error("Failed to put account", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	return nil
}

func (a *Account) unmarshal(b []byte) error {
	if err := json.Unmarshal(b, &a); err != nil {
		pluto.Log.Error("Can not unmarshal account", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to find account",
		}
	}
	return nil
}
