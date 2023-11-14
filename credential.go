package pluto

type Credential interface {
	Validate(Identifier) (bool, error)
}

type OutComingCredential struct {
	Token string `json:"token"`
}

func (c OutComingCredential) Validate(Identifier) (bool, error) {
	return true, nil
}
