package panel

import (
	"fmt"
	"os"
	"pluto/panel/account"
	"time"
)

func init() {
	accounts, err := account.GetStorage().All()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Get accounts: %v\n", err)
		os.Exit(1)
	}

	if len(accounts) > 0 {
		return
	}

	if err := (&account.Account{
		Email:    "admin",
		Password: account.MustNewPassword([]byte("admin")),
		SavedAt:  time.Now(),
	}).Save(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create the default admin account: %v\n", err)
		os.Exit(1)
	}
}
