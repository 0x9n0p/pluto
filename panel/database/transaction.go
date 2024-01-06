package database

import (
	"net/http"
	"pluto"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type Transaction struct {
	*bbolt.Tx
}

func (t *Transaction) CommitOrRollback() error {
	if err := t.Commit(); err != nil {
		if err2 := t.Rollback(); err2 != nil {
			return err2
		}
		return err
	}

	return nil
}

func (t *Transaction) Commit() error {
	if !t.Writable() {
		return nil
	}

	if err := t.Tx.Commit(); err != nil {
		pluto.Log.Error("Failed to commit transaction", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "An internal server error occurred",
		}
	}
	return nil
}

func (t *Transaction) Rollback() error {
	if err := t.Tx.Rollback(); err != nil {
		pluto.Log.Error("Failed to rollback transaction", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "An internal server error occurred",
		}
	}
	return nil
}
