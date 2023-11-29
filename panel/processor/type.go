package processor

const (
	TypeText      = "Text"
	TypeNumeric   = "Numeric"
	TypeList      = "List" // List of Value
	TypeProcessor = "Processor"
)

type ValueDescriptor struct {
	Type string `json:"type"`
	//Value    any    `json:"value"`
	Default  any  `json:"default,omitempty"`
	Required bool `json:"required"`
}

type Value struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}
