package pluto

// Note:
//   1. The returning errors should be pluto.Error
//   2. Extensions are not going to resolve conflicts, they will change pipelines/processors if already exists

type Extension interface {
	Install() error
	Uninstall() error
}
