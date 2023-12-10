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
			e, ok := v.(error)
			if !ok {
				pluto.Log.Error("The value of recovered panic is not an error", zap.Any("value", v))
			}

			err = &pluto.Error{
				HTTPCode: http.StatusBadRequest,
				Message:  fmt.Sprintf("%v", e),
			}
		}
	}()

	if err := p.validateArguments(); err != nil {
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

func (p *Processor) validateArguments() error {
	descriptor, found := GetDescriptor(p.Name)
	if !found {
		return &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Processor descriptor (%s) not found", p.Name),
		}
	}

	for _, valueDescriptor := range descriptor.Arguments {
		argument, index := pluto.MayFind(valueDescriptor.Name, p.Arguments...)
		if index == -1 || argument.Value == nil {
			if valueDescriptor.Required {
				return &pluto.Error{
					HTTPCode: http.StatusBadRequest,
					Message:  fmt.Sprintf("Argument (%s) for processor (%s) is required", valueDescriptor.Name, p.Name),
				}
			} else {
				argument = pluto.Value{
					Name:  valueDescriptor.Name,
					Type:  valueDescriptor.Type,
					Value: valueDescriptor.Default,
				}
				index = len(p.Arguments)
				p.Arguments = append(p.Arguments, argument)
			}
		}

		{
			if valueDescriptor.ValueValidator == nil {
				valueDescriptor.ValueValidator = pluto.DefaultValueValidator
			}

			if err := valueDescriptor.ValueValidator(argument, valueDescriptor); err != nil {
				return &pluto.Error{
					HTTPCode: http.StatusBadRequest,
					Message:  fmt.Sprintf("%v", err),
				}
			}
		}

		argument.ValueParser = GetParserByType(argument.Type)
		p.Arguments[index] = argument
	}

	return nil
}

func GetParserByType(t string) func(any) any {
	switch t {
	case pluto.TypeNumeric:
		return func(v any) any {
			return int(v.(float64))
		}
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
				// The caller of this parser recovers this error.
				panic(err)
			}

			return p
		}
	default:
		return pluto.NoParserRequired
	}
}
