package account

import (
	"net/http"
	"pluto"
	"pluto/pkg/random"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

/*
	TODO
		Implement access tokens and refresh tokens
*/

var (
	JWTExpiration = time.Hour * 24
	JWTSecretKey  = []byte(random.String(32))
)

type JsonWebToken struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func NewJsonWebToken(email string) JsonWebToken {
	return JsonWebToken{
		Email: email,
	}
}

func (p *JsonWebToken) Create() (err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": p.Email,
			"exp":   time.Now().Add(JWTExpiration).Unix(),
		},
	)

	p.Token, err = token.SignedString(JWTSecretKey)
	if err != nil {
		pluto.Log.Error("Failed to generate json web token", zap.String("email", p.Email), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Internal server error",
		}
	}

	return
}
