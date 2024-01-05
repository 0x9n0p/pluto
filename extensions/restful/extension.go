package restful

import (
	"pluto"
	"pluto/extensions"
)

const ExtensionID_V1 = "restful-v1"

func init() {
	extensions.Extensions[ExtensionID_V1] = V1
}

var V1 = &Restful{
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
	install   func() error
	uninstall func() error
}

func (r *Restful) Install() error {
	return r.install()
}

func (r *Restful) Uninstall() error {
	return r.uninstall()
}
