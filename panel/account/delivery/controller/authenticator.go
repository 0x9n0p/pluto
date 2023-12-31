package controller

import (
	"net/http"
	"pluto/panel/account"
	"pluto/panel/account/delivery"
	"pluto/panel/pkg/wrapper"
	"time"
)

type Authenticator struct {
	AuthenticationMethod int `json:"authenticator_method"`

	Email    string `json:"email"`
	Password string `json:"password"`

	SignedMessage string `json:"signed_message"`

	Authenticator account.Authenticator `json:"-"`
}

func (c *Authenticator) Exec(w wrapper.ResponseWriter) (err error) {
	if err := c.Authenticator.Authenticate(); err != nil {
		return delivery.WriteError(err, w)
	}

	jwt := account.NewJsonWebToken(c.Email)
	if err := jwt.Create(); err != nil {
		return delivery.WriteError(err, w)
	}

	{
		expires := time.Now().UTC().Add(account.JWTExpiration)

		w.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    jwt.Token,
			Expires:  expires,
			Secure:   true,
			HttpOnly: true,
		})

		w.SetCookie(&http.Cookie{
			Name:     "email",
			Value:    jwt.Email,
			Expires:  expires,
			Secure:   true,
			HttpOnly: true,
		})
	}

	return w.JSON(http.StatusOK, jwt)
}
