package pluto

const (
	TypeText              = "Text"
	TypeNumeric           = "Numeric"
	TypeList              = "List" // List of Value
	TypeProcessor         = "Processor"
	TypeBytes             = "Bytes"
	TypeInternalInterface = "InternalInterface"
)

var (
	/*
		These parsers can panic.
	*/

	NoParserRequired    = func(v any) any { return v }
	TypeProcessorParser = func(v any) any {
		m := v.(map[string]any)
		name := m["name"].(string)
		arguments := m["arguments"].([]Value)
		return PredefinedProcessors[name](arguments)
	}
)

type ValueDescriptor struct {
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	Required       bool              `json:"required"`
	Default        any               `json:"default,omitempty"`
	ValueValidator func(Value) error `json:"-"`
	ValueParser    func(any) any     `json:"-"`
}

func (v ValueDescriptor) Comparable() any {
	return v.Name
}

type Value struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Value       any           `json:"value"`
	ValueParser func(any) any `json:"-"`
}

func (v Value) Get() any {
	return v.ValueParser(v.Value)
}

func (v Value) Comparable() any {
	return v.Name
}
