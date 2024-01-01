package delivery

import (
	"net/http"
	"pluto"
	"pluto/panel/account"
	"pluto/panel/account/delivery/controller"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := pluto.FindHTTPHost(pluto.PanelSubdomain).Group("/api/v1").Group("/account")
	authenticated := v1.Group("", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.POST("/login",
		wrapper.New[controller.Authenticator](
			func(c controller.Authenticator, w wrapper.ResponseWriter) error {
				if err := FactoryAuthenticator(&c); err != nil {
					return w.Error(err.(wrapper.HTTPResponseError))
				}

				return c.Exec(w)
			},
		).Handle(),
	)

	/*
		TODO
			Use access token and refresh token to improve logout with short token expiration.
	*/

	authenticated.POST("/logout",
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
				return wrapper.WriteError(err, w)
			}

			return w.JSON(http.StatusOK, a)
		}).Handle(),
	)

	authenticated.PATCH("/password",
		wrapper.New[controller.ChangePassword](func(c controller.ChangePassword, w wrapper.ResponseWriter) error {
			return c.Exec(w)
		}).Handle(),
	)
}
