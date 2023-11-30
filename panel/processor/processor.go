package processor

import (
	"fmt"
	"net/http"
	"pluto"
)

type Descriptor struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Icon        string                  `json:"icon"`
	Arguments   []pluto.ValueDescriptor `json:"arguments"`
	Input       []pluto.ValueDescriptor `json:"input"`
	Output      []pluto.ValueDescriptor `json:"output"`
	Category    string                  `json:"category"`
}

type Processor struct {
	Name      string        `json:"name"`
	Arguments []pluto.Value `json:"arguments"`
}

func (p *Processor) Create() (pluto.Processor, error) {
	descriptor, found := GetDescriptor(p.Name)
	if !found {
		return nil, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Processor descriptor (%s) not found", p.Name),
		}
	}

	if err := p.validateArguments(descriptor.Arguments); err != nil {
		return nil, err
	}

	creator, found := pluto.PredefinedProcessors[p.Name]
	if !found {
		return nil, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Processor creator (%s) not found", p.Name),
		}
	}

	return creator(p.Arguments), nil
}

func (p *Processor) validateArguments(descriptors []pluto.ValueDescriptor) error {
	for _, descriptor := range descriptors {
		argument, found := pluto.MayFind(pluto.Value{Name: descriptor.Name}, p.Arguments...)
		if !found {
			return &pluto.Error{
				HTTPCode: http.StatusBadRequest,
				Message:  fmt.Sprintf("Argument (%s) is required", descriptor.Name),
			}
		}

		if descriptor.Type != argument.Type {
			return &pluto.Error{
				HTTPCode: http.StatusBadRequest,
				Message:  fmt.Sprintf("Type (%s) is not the required type (%s)", argument.Type, descriptor.Type),
			}
		}

		if err := descriptor.ValueValidator(argument); err != nil {
			return err
		}
	}

	return nil
}
