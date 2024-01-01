package delivery

import (
	"pluto/panel/account"

	echojwt "github.com/labstack/echo-jwt/v4"
)

var DefaultJWTConfig = echojwt.Config{
	SigningKey:  account.JWTSecretKey,
	TokenLookup: "header:Authorization:Bearer ,cookie:token",
}
