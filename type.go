package pluto

const (
	TypeText              = "Text"
	TypeNumeric           = "Numeric"
	TypeList              = "List" // List of Value
	TypeProcessor         = "Processor"
	TypeBytes             = "Bytes"
	TypeInternalInterface = "InternalInterface"
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
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (v Value) Comparable() any {
	return v.Name
}
