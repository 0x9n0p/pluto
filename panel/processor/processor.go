package processor

import (
	"fmt"
	"net/http"
	"pluto"

	"go.uber.org/zap"
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
	Name      string        `json:"name" query:"name"`
	Arguments []pluto.Value `json:"arguments"`
}

func (p *Processor) Create() (processor pluto.Processor, err error) {
	defer func() {
		if v := recover(); v != nil {
			var ok bool
			err, ok = v.(error)
			if !ok {
				pluto.Log.Error("The value of recovered panic is not an error", zap.Any("value", v))
			}
		}
	}()

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

	processor, err = creator(p.Arguments)
	return processor, err
}

func (p *Processor) validateArguments(descriptors []pluto.ValueDescriptor) error {
	for _, descriptor := range descriptors {
		argument, index := pluto.MayFind(descriptor.Name, p.Arguments...)
		if index == -1 {
			return &pluto.Error{
				HTTPCode: http.StatusBadRequest,
				Message:  fmt.Sprintf("Argument (%s) for processor (%s) is required", descriptor.Name, p.Name),
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

		argument.ValueParser = GetParserByType(argument.Type)
		p.Arguments[index] = argument
	}

	return nil
}

func GetParserByType(t string) func(any) any {
	switch t {
	case pluto.TypeProcessor:
		return func(v any) any {
			m := v.(map[string]any)

			processor := Processor{
				Name:      m["name"].(string),
				Arguments: []pluto.Value{},
			}

			for _, varg := range m["arguments"].([]any) {
				processor.Arguments = append(processor.Arguments, pluto.ValueFromMap(varg.(map[string]any)))
			}

			p, err := processor.Create()
			if err != nil {
				panic(err)
			}

			return p
		}
	default:
		return pluto.NoParserRequired
	}
}
