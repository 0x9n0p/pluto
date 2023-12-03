package pluto

const (
	TypeText    = "Text"
	TypeNumeric = "Numeric"
	TypeList    = "List" // List of Value

	/*
		TypeProcessor is not a struct.
		It must be a map[string]any and has two fields 'string name' and '[]any arguments'.
	*/
	TypeProcessor         = "Processor"
	TypeBytes             = "Bytes"
	TypeInternalInterface = "InternalInterface"
)

var (
	/*
		Parsers can panic.
	*/

	NoParserRequired = func(v any) any { return v }
)

type ValueDescriptor struct {
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	Required       bool              `json:"required"`
	Default        any               `json:"default,omitempty"`
	ValueValidator func(Value) error `json:"-"`
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

// ValueFromMap can panic.
func ValueFromMap(m map[string]any) Value {
	return Value{
		Name:  m["name"].(string),
		Type:  m["type"].(string),
		Value: m["value"].(string),
	}
}

func (v Value) Get() any {
	return v.ValueParser(v.Value)
}

func (v Value) Comparable() any {
	return v.Name
}
