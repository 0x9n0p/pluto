package processor

import (
	"io"
	"net/http"
	"pluto"
)

const (
	Category_InputOutpt = "InputOutput"
)

var Processors = []Descriptor{
	{
		Name:        pluto.ProcessorName_WriteToInputOutput,
		Description: "Write to Input/Output interfaces directly",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "io_interface",
				Type:     pluto.TypeInternalInterface,
				Required: true,
				ValueValidator: func(arg pluto.Value) error {
					_, ok := arg.Value.(io.Writer)
					if !ok {
						return &pluto.Error{
							HTTPCode: http.StatusBadRequest,
							Message:  "Value of io_interface is not an io.Writer",
						}
					}
					return nil
				},
			},
		},
		Input: []pluto.ValueDescriptor{
			/*
				The processable.body is Processable.GetBody()
			*/
			{
				Name:     "processable.body",
				Type:     pluto.TypeBytes,
				Required: true,
			},
		},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_InputOutpt,
	},
}

func GetDescriptor(name string) (Descriptor, bool) {
	for _, processor := range Processors {
		if name == processor.Name {
			return processor, true
		}
	}
	return Descriptor{}, false
}
