package account

import (
	"pluto"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Password []byte

func NewPassword(plain []byte) (Password, error) {
	return bcrypt.GenerateFromPassword(plain, 14)
}

func MustNewPassword(plain []byte) Password {
	pass, err := NewPassword(plain)
	if err != nil {
		pluto.Log.Debug("Failed to hash the password",
			zap.String("password", string(plain)), // NOTE: Do not use debug mode in production.
			zap.Error(err),
		)
		return Password{}
	}

	return pass
}

func (p *Password) Compare(plain []byte) bool {
	return bcrypt.CompareHashAndPassword(*p, plain) == nil
}

func (p *Password) String() string {
	return string(*p)
}
