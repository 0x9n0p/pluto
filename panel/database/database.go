package database

import (
	"fmt"
	"net/http"
	"os"
	"pluto"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

var database *Database

func Get() *Database {
	return database
}

func setup() {
	db, err := bbolt.Open(Env.DatabasePath, 0600, nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Open database: %s: %v\n", Env.DatabasePath, err)
		os.Exit(1)
	}

	database = &Database{db}
}

type Database struct {
	*bbolt.DB
}

func (s *Database) NewTransaction(write bool) (*Transaction, error) {
	begin, err := s.DB.Begin(write)
	if err != nil {
		pluto.Log.Error("Failed to create transaction", zap.Error(err))
		return nil, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to begin operation",
		}
	}
	return &Transaction{begin}, nil
}
