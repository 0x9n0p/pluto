package restful

import (
	"pluto"
)

var V1 = Restful{
	ExtensionDescriptor: pluto.ExtensionDescriptor{
		ID:         "restful-v1",
		Name:       "Restful",
		Processors: []string{},
		Pipelines:  []string{},
	},
}
