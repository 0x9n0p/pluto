package pluto

type Extension interface {
	Install() error
	Uninstall() error
	GetExtensionDescriptor() ExtensionDescriptor
}

type ExtensionDescriptor struct {
	// ID such as extension_name-v1.2.3
	ID   string
	Name string

	/*
		Below pipelines and processors are added/deleted during the installation/uninstallation process.
	*/

	Processors []string
	Pipelines  []string
}
