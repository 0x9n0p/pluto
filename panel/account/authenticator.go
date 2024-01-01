package account

type Authenticator interface {
	Authenticate() error
}
