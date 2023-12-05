package delivery

import (
	"errors"
	"net/http"
	"pluto"
	"pluto/panel/auth"
	"pluto/panel/pkg/wrapper"
)

func init() {
	WriteError := func(err error, writer wrapper.ResponseWriter) error {
		var perr *pluto.Error
		if errors.As(err, &perr) {
			return writer.JSON(perr.HTTPCode, perr)
		}
		return writer.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	v1 := pluto.FindHTTPHost("panel").Group("/api/v1")

	v1.POST("/auth",
		wrapper.New[auth.PasswordAuthenticator](func(p auth.PasswordAuthenticator, writer wrapper.ResponseWriter) error {
			if err := p.Authenticate(); err != nil {
				return WriteError(err, writer)
			}

			jwt := auth.NewJsonWebToken(p.Email)
			if err := jwt.Create(); err != nil {
				return WriteError(err, writer)
			}

			return writer.JSON(http.StatusOK, jwt)
		}).Handle(),
	)
}
