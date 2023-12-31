package controller

import (
	"net/http"
	"pluto/panel/account"
	"pluto/panel/pkg/wrapper"
)

type ChangePassword struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (c *ChangePassword) Exec(w wrapper.ResponseWriter) (err error) {
	claims, err := wrapper.GetJWTClaims(w)
	if err != nil {
		return err
	}

	a, err := account.Find(claims["email"].(string))
	if err != nil {
		return wrapper.WriteError(err, w)
	}

	if err := a.ChangePassword([]byte(c.OldPassword), []byte(c.NewPassword)); err != nil {
		return wrapper.WriteError(err, w)
	}

	return w.NoContent(http.StatusOK)
}
