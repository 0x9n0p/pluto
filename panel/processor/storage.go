package processor

import (
	"fmt"
	"io"
	"net/http"
	"pluto"
	"strings"
)

const (
	Category_Flow          = "Flow"
	Category_InputOutpt    = "InputOutput"
	Category_Communication = "Communication"
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
				ValueValidator: func(arg pluto.Value, _ pluto.ValueDescriptor) error {
					_, ok := arg.Value.(io.Writer)
					if !ok {
						return &pluto.Error{
							HTTPCode: http.StatusBadRequest,
							Message:  "Value of (io_interface) is not an io.Writer",
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
	{
		Name:        "Execute processor and join the result",
		Description: "Execute processor and join the result ..",
		Icon:        "https://..",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "Processor",
				Type:     pluto.TypeProcessor,
				Required: true,
				ValueValidator: func(arg pluto.Value, _ pluto.ValueDescriptor) (err error) {
					m, ok := arg.Value.(map[string]any)
					if !ok {
						return &pluto.Error{
							HTTPCode: http.StatusBadRequest,
							Message:  "Value of (Processor) is not a processor",
						}
					}

					defer func() {
						if v := recover(); v != nil {
							err = &pluto.Error{
								HTTPCode: http.StatusBadRequest,
								Message:  fmt.Sprintf("Missing fields or incorrect types: %s", v),
							}
						}

					}()

					_ = m["name"].(string)
					_ = m["arguments"].([]any)

					return
				},
			},
		},
		Input: []pluto.ValueDescriptor{
			/*
				Input of the inner processor.
			*/
		},
		Output: []pluto.ValueDescriptor{
			/*
				Output of the inner processor.
			*/
		},
		Category: Category_Flow,
	},
	{
		Name:        pluto.ProcessorName_CreateChannel,
		Description: "Creates a channel",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "Name",
				Type:     pluto.TypeText,
				Required: true,
			},
			{
				Name:    "Length",
				Type:    pluto.TypeNumeric,
				Default: 10,
			},
		},
		Input: []pluto.ValueDescriptor{},
		Output: []pluto.ValueDescriptor{
			{
				Name:     "processable.body",
				Type:     pluto.TypeChannel,
				Required: true,
			},
		},
		Category: Category_Communication,
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

type DescriptorFinder struct {
	Name string `query:"name"`
}

func (f *DescriptorFinder) Find() []Descriptor {
	if f.Name == "" {
		return Processors
	}

	found := make([]Descriptor, 0)
	for _, descriptor := range Processors {
		if strings.Contains(strings.ToLower(descriptor.Name), strings.ToLower(f.Name)) {
			found = append(found, descriptor)
		}
	}
	return found
}
