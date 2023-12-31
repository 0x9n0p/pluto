package delivery

import (
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/account"
	"pluto/panel/account/delivery/controller"
)

const (
	AuthenticationMethod_Password = iota
	AuthenticationMethod_PublicKey
)

func FactoryAuthenticator(c *controller.Authenticator) (err error) {
	switch c.AuthenticationMethod {
	case AuthenticationMethod_Password:
		c.Authenticator = &account.PasswordAuthenticator{
			Email:    c.Email,
			Password: c.Password,
		}
		break
	case AuthenticationMethod_PublicKey:
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
