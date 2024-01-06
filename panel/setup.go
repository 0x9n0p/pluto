package panel

import (
	"fmt"
	"os"
	"pluto/panel/account"
	"pluto/panel/database"
	"time"
)

func init() {
	tx, err := database.Get().NewTransaction(true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Setup panel: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		if err := tx.CommitOrRollback(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Setup panel: %v\n", err)
			os.Exit(1)
		}
	}()

	accounts, err := account.All(tx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Setup panel: %v\n", err)
		os.Exit(1)
	}

	if len(accounts) > 0 {
		return
	}

	if err := (&account.Account{
		Email:       "admin",
		Password:    account.MustNewPassword([]byte("admin")),
		SavedAt:     time.Now(),
		Transaction: tx,
	}).Save(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create the default admin account: %v\n", err)
		os.Exit(1)
	}
}
