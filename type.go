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

func (v ValueDescriptor) Compare(c Comparable) (equal bool) {
	t, ok := c.(ValueDescriptor)
	if !ok {
		return false
	}
	return v.Name == t.Name
}

type Value struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (v Value) Compare(c Comparable) bool {
	t, ok := c.(Value)
	if !ok {
		return false
	}
	return v.Name == t.Name
}
