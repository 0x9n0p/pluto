package pluto

import "fmt"

const ProcessorName_NumberValidator = "NUMBER_VALIDATOR"

func init() {
	PredefinedProcessors[ProcessorName_NumberValidator] = func(args []Value) (p Processor, err error) {
		defer CreatorPanicHandler(ProcessorName_NumberValidator, &err)()
		return NumberValidator{
			Name:    Find("name", args...).Get().(string),
			Minimum: Find("minimum", args...).Get().(int),
			Maximum: Find("maximum", args...).Get().(int),
		}, err
	}
}

type NumberValidator struct {
	Name    string
	Minimum int
	Maximum int

	// TODO
	//  Required or a default value
}

func (p NumberValidator) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_NumberValidator},
		})
		return processable, false
	}

	v, found := appendable[p.Name]
	if !found {
		ApplicationLogger.Debug(ApplicationLog{
			Message: fmt.Sprintf("Input (%s) is required", p.Name),
			Extra:   map[string]any{"issuer": ProcessorName_NumberValidator},
		})
		return processable, false
	}

	number, ok := v.(Number)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: fmt.Sprintf("Input (%s) is not a number", p.Name),
			Extra: map[string]any{
				"issuer": ProcessorName_NumberValidator,
				"value":  v,
			},
		})
		return processable, false
	}

	converted, err := number.Int64()
	if err != nil {
		ApplicationLogger.Debug(ApplicationLog{
			Message: fmt.Sprintf("Input (%s) is not a number", p.Name),
			Extra: map[string]any{
				"issuer": ProcessorName_NumberValidator,
				"value":  number,
			},
		})
		return processable, false
	}

	if int(converted) < p.Minimum {
		ApplicationLogger.Debug(ApplicationLog{
			Message: fmt.Sprintf("Input (%s) is less than the minimum", p.Name),
			Extra: map[string]any{
				"issuer":  ProcessorName_NumberValidator,
				"value":   number,
				"minimum": p.Minimum,
			},
		})
		return processable, false
	}

	if int(converted) > p.Maximum {
		ApplicationLogger.Debug(ApplicationLog{
			Message: fmt.Sprintf("Input (%s) is bigger than the maximum", p.Name),
			Extra: map[string]any{
				"issuer":  ProcessorName_NumberValidator,
				"value":   number,
				"maximum": p.Maximum,
			},
		})
		return processable, false
	}

	return processable, true
}
