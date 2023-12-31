package delivery

import (
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/account"
	"pluto/panel/account/delivery/controller"
)

const (
	PasswordAuthenticationMethod = iota
	PublicKeyAuthenticationMethod
)

func FactoryAuthenticator(c *controller.Authenticator) (err error) {
	switch c.AuthenticationMethod {
	case PasswordAuthenticationMethod:
		c.Authenticator = &account.PasswordAuthenticator{
			Email:    c.Email,
			Password: account.MustNewPassword([]byte(c.Password)),
		}
		break
	case PublicKeyAuthenticationMethod:
		c.Authenticator = &account.PublicKeyAuthenticator{
			SignedMessage: c.SignedMessage,
		}
		break
	default:
		return &pluto.Error{
			HTTPCode: http.StatusBadRequest,
			Message:  fmt.Sprintf("Authentication method (%d) not found", c.AuthenticationMethod),
		}
	}

	return
}
