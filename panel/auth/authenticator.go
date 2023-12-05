package auth

type Authenticator interface {
	Authenticate() error
}
