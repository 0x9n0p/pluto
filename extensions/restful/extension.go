package restful

import (
	"pluto"
)

var V1 = Restful{
	ExtensionDescriptor: pluto.ExtensionDescriptor{
		ID:   "restful-v1",
		Name: "Restful",
		Processors: []string{
			ProcessorName_WriteResponse,
		},
		Pipelines: []string{},
	},
	install: func() error {
		// V0.install() instead of struct composition.
		pluto.PredefinedProcessors[ProcessorName_WriteResponse] = creator_WriteResponse
		return nil
	},
	uninstall: func() error {
		delete(pluto.PredefinedProcessors, ProcessorName_WriteResponse)
		return nil
	},
}

type Restful struct {
	pluto.ExtensionDescriptor
	install   func() error
	uninstall func() error
}

func (r *Restful) Install() error {
	return r.install()
}

func (r *Restful) Uninstall() error {
	return r.uninstall()
}

func (r *Restful) GetExtensionDescriptor() pluto.ExtensionDescriptor {
	return r.ExtensionDescriptor
}
