package account

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"pluto"
	"time"

	"go.uber.org/zap"
)

type Account struct {
	Email     string    `json:"email"`
	Password  Password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	SavedAt   time.Time `json:"saved_at"`
}

func Find(email string) (a Account, err error) {
	f, err := os.OpenFile(StorageABSPath(email), os.O_RDWR, 0644)
	if err != nil {
		pluto.Log.Debug("Account not found or a system error occurred", zap.String("path", StorageABSPath(email)), zap.Error(err))
		return Account{}, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  "Can not find account",
		}
	}

	b, err := io.ReadAll(f)
	if err != nil {
		pluto.Log.Error("Can not read account file", zap.Error(err))
		return Account{}, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Can not read account",
		}
	}

	if err := json.Unmarshal(b, &a); err != nil {
		pluto.Log.Error("Can not unmarshal account", zap.Error(err))
		return Account{}, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Can not read account",
		}
	}

	return
}

// ChangeEmail
// TODO
func (a *Account) ChangeEmail(new string) error {
	return &pluto.Error{
		HTTPCode: http.StatusNotImplemented,
		Message:  "Not Implemented",
	}
}

func (a *Account) ChangePassword(old, new []byte) error {
	oldPass, err := NewPassword(old)
	if err != nil {
		return &pluto.Error{
			HTTPCode: http.StatusBadRequest,
			Message:  "Make sure you have entered the old password in the correct format",
		}
	}

	newPass, err := NewPassword(new)
	if err != nil {
		return &pluto.Error{
			HTTPCode: http.StatusBadRequest,
			Message:  "Make sure you have entered the new password in the correct format",
		}
	}

	if !a.Password.Compare(oldPass) {
		return &pluto.Error{
			HTTPCode: http.StatusUnauthorized,
			Message:  "The entered old password is incorrect",
		}
	}

	a.Password = newPass
	return a.save()
}

func (a *Account) save() error {
	b, err := json.Marshal(a)
	if err != nil {
		pluto.Log.Error("Failed to marshal the account", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	tmpABSPath := StorageABSPath(a.Email) + "-tmp"

	f, err := os.OpenFile(tmpABSPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		pluto.Log.Error("Failed to create temporary account", zap.String("path", tmpABSPath), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	if _, err := f.Write(b); err != nil {
		pluto.Log.Error("Failed to write to temporary account", zap.String("path", tmpABSPath), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	if err := os.Rename(tmpABSPath, StorageABSPath(a.Email)); err != nil {
		pluto.Log.Error("Failed to move temporary account",
			zap.String("src", tmpABSPath),
			zap.String("dst", StorageABSPath(a.Email)),
			zap.Error(err),
		)
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save account",
		}
	}

	return nil
}

func StorageABSPath(email string) string {
	return filepath.Join(Env.StoragePath, email)
}
