package controller

import (
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/database"
	"pluto/panel/extension"
	"pluto/panel/pkg/wrapper"
)

type InstallExtension struct {
	ID string `param:"extension_id" validate:"required"`
}

func (c *InstallExtension) Exec(w wrapper.ResponseWriter) (err error) {
	descriptor, found := extension.FindDescriptor(c.ID)
	if !found {
		return &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("No extension found by id (%s)", c.ID),
		}
	}

	transaction, err := database.Get().NewTransaction(true)
	if err != nil {
		return wrapper.WriteError(err, w)
	}

	e := extension.Extension{
		Descriptor:  descriptor,
		Transaction: transaction,
	}

	if err := e.Install(); err != nil {
		_ = transaction.Rollback()
		return wrapper.WriteError(err, w)
	}

	_ = transaction.CommitOrRollback()
	return w.JSON(http.StatusOK, e)
}
