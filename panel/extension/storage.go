package extension

import (
	"fmt"
	"net/http"
	"pluto"
	"pluto/extensions"
	"pluto/extensions/restful"
)

var storage = Storage{
	Descriptors: []Descriptor{
		{
			ID:   restful.ExtensionID_V1,
			Name: "Restful",
			Processors: []string{
				restful.ProcessorName_WriteResponse,
			},
			Pipelines: []string{},
		},
	},
}

func GetStorage() Storage {
	return storage
}

type Storage struct {
	Descriptors []Descriptor
}

func (s *Storage) Get(id string) (pluto.Extension, error) {
	e, found := extensions.Extensions[id]
	if !found {
		return nil, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Extension (%s) not found", id),
		}
	}
	return e, nil
}
