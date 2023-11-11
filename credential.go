package pluto

type Credential interface {
	Validate() (bool, error)
}
