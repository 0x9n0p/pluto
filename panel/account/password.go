package account

import "golang.org/x/crypto/bcrypt"

type Password []byte

func NewPassword(plain []byte) (Password, error) {
	return bcrypt.GenerateFromPassword(plain, 14)
}

func (p *Password) Compare(hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, *p) == nil
}

func (p *Password) String() string {
	return string(*p)
}
