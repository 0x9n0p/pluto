package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pluto"
	"pluto/panel/database"

	"go.uber.org/zap"
)

var bucket = []byte("account")

func init() {
	transaction, err := database.Get().NewTransaction(true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
		os.Exit(1)
	}

	defer func() {
		if err := transaction.CommitOrRollback(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
			os.Exit(1)
		}
	}()

	if _, err := transaction.CreateBucketIfNotExists(bucket); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
		os.Exit(1)
	}
}

func Find(tx *database.Transaction, email string) (a Account, err error) {
	b := tx.Bucket(bucket).Get([]byte(email))
	if b == nil {
		return Account{}, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Account (%s) not found", email),
		}
	}

	if err := json.Unmarshal(b, &a); err != nil {
		pluto.Log.Error("Can not unmarshal account", zap.Error(err))
		return Account{}, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to find account",
		}
	}

	a.Transaction = tx

	return
}

func All(tx *database.Transaction) ([]Account, error) {
	accounts := []Account{} // Do not replace it with a nil slice declaration
	return accounts, tx.Bucket(bucket).ForEach(func(k, v []byte) error {
		var a Account
		if err := a.unmarshal(v); err != nil {
			return err
		}
		a.Transaction = tx
		return nil
	})
}
