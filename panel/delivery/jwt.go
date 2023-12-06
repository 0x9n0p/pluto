package delivery

import (
	"pluto/panel/auth"

	echojwt "github.com/labstack/echo-jwt/v4"
)

var DefaultJWTConfig = echojwt.Config{
	SigningKey:  auth.JWTSecretKey,
	TokenLookup: "header:Authorization:Bearer ,cookie:token",
}
