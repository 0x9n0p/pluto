package delivery

import (
	"errors"
	"net/http"
	"pluto"
	"pluto/panel/account"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := pluto.FindHTTPHost(pluto.PanelSubdomain).Group("/api/v1").Group("/account")
	authenticated := v1.Group("", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.POST("/login",
		wrapper.New[account.PasswordAuthenticator](func(p account.PasswordAuthenticator, writer wrapper.ResponseWriter) error {
			if err := p.Authenticate(); err != nil {
				return WriteError(err, writer)
			}

			jwt := account.NewJsonWebToken(p.Email)
			if err := jwt.Create(); err != nil {
				return WriteError(err, writer)
			}

			{
				expires := time.Now().UTC().Add(account.JWTExpiration)

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

	authenticated.Group("", echojwt.WithConfig(delivery.DefaultJWTConfig)).POST("/logout",
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

	authenticated.PATCH("/password",
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
