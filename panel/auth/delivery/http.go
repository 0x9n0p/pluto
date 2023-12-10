package delivery

import (
	"errors"
	"net/http"
	"pluto"
	"pluto/panel/auth"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
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
				expires := time.Now().UTC().Add(auth.JWTExpiration)

				writer.SetCookie(&http.Cookie{
					Name:     "token",
					Value:    jwt.Token,
					Expires:  expires,
					Secure:   true,
					HttpOnly: true,
				})

				writer.SetCookie(&http.Cookie{
					Name:     "email",
					Value:    jwt.Email,
					Expires:  expires,
					Secure:   true,
					HttpOnly: true,
				})
			}

			return writer.JSON(http.StatusOK, jwt)
		}).Handle(),
	)

	/*
		TODO
			Use access token and refresh token to improve logout with short token expiration.
	*/

	v1.Group("", echojwt.WithConfig(delivery.DefaultJWTConfig)).POST("/logout",
		wrapper.New[wrapper.EmptyRequest](func(_ wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			writer.SetCookie(&http.Cookie{
				Name:    "token",
				Expires: time.Now().UTC(),
			})

			writer.SetCookie(&http.Cookie{
				Name:    "email",
				Expires: time.Now().UTC(),
			})

			return writer.NoContent(http.StatusOK)
		}).Handle(),
	)
}
