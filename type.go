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
	Type           string            `json:"type"`
	Required       bool              `json:"required"`
	Default        any               `json:"default,omitempty"`
	ValueValidator func(Value) error `json:"-"`
}

type Value struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}
