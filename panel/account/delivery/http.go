package delivery

import (
	"errors"
	"net/http"
	"pluto"
	"pluto/panel/account"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := pluto.FindHTTPHost(pluto.PanelSubdomain).Group("/api/v1")
	authenticated := v1.Group("/account", echojwt.WithConfig(delivery.DefaultJWTConfig))

	authenticated.GET("",
		wrapper.New[wrapper.EmptyRequest](func(_ wrapper.EmptyRequest, w wrapper.ResponseWriter) error {
			claims, err := wrapper.GetJWTClaims(w)
			if err != nil {
				return err
			}

			a, err := account.Find(claims["email"].(string))
			if err != nil {
				return WriteError(err, w)
			}

			return w.JSON(http.StatusOK, a)
		}).Handle(),
	)

	authenticated.POST("/password",
		wrapper.New[ChangePasswordController](func(c ChangePasswordController, w wrapper.ResponseWriter) error {
			return c.Exec(w)
		}).Handle(),
	)
}

func WriteError(err error, writer wrapper.ResponseWriter) error {
	var perr *pluto.Error
	if errors.As(err, &perr) {
		return writer.JSON(perr.HTTPCode, perr)
	}
	return writer.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
}
