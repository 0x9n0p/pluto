package account

import (
	"fmt"
	"net/http"
	"os"
	"pluto"

	"go.uber.org/zap"
)

var storage *Storage

func GetStorage() *Storage {
	return storage
}

func init() {
	if err := os.MkdirAll(Env.AccountsPath, 0644); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create storage directory: %s\n", Env.AccountsPath)
		os.Exit(1)
	}

	storage = &Storage{Env.AccountsPath}
}

type Storage struct {
	Path string
}

func (s *Storage) All() ([]Account, error) {
	files, err := os.ReadDir(Env.AccountsPath)
	if err != nil {
		pluto.Log.Debug("Failed to get accounts", zap.String("path", Env.AccountsPath), zap.Error(err))
		return nil, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to get accounts",
		}
	}

	accounts := []Account{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		a, err := Find(file.Name())
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}
