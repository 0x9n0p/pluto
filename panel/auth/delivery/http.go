package delivery

import (
	"errors"
	"net/http"
	"pluto"
	"pluto/panel/auth"
	"pluto/panel/pkg/wrapper"
	"time"
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

			{
				expires := time.Now().Add(auth.JWTExpiration)

				writer.SetCookie(&http.Cookie{
					Name:     "token",
					Value:    jwt.Token,
					Expires:  expires,
					Secure:   false,
					HttpOnly: true,
				})

				writer.SetCookie(&http.Cookie{
					Name:     "email",
					Value:    jwt.Email,
					Expires:  expires,
					Secure:   false,
					HttpOnly: true,
				})
			}

			return writer.JSON(http.StatusOK, jwt)
		}).Handle(),
	)
}
