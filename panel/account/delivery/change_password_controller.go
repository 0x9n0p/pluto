package delivery

import (
	"net/http"
	"pluto/panel/account"
	"pluto/panel/pkg/wrapper"
)

type ChangePasswordController struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (c *ChangePasswordController) Exec(w wrapper.ResponseWriter) (err error) {
	claims, err := wrapper.GetJWTClaims(w)
	if err != nil {
		return err
	}

	a, err := account.Find(claims["email"].(string))
	if err != nil {
		return WriteError(err, w)
	}

	if err := a.ChangePassword([]byte(c.OldPassword), []byte(c.NewPassword)); err != nil {
		return WriteError(err, w)
	}

	return w.NoContent(http.StatusOK)
}
