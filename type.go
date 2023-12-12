package pluto

import "fmt"

const (
	TypeBoolean = "Boolean"
	TypeText    = "Text"
	TypeNumeric = "Numeric"
	TypeList    = "List" // List of Value

	/*
		TypeProcessor is not a struct.
		It must be a map[string]any and has two fields 'string name' and '[]any arguments'.
	*/
	TypeProcessor         = "Processor"
	TypeChannel           = "Channel"
	TypeIdentifier        = "Identifier"
	TypeBytes             = "Bytes"
	TypeInternalInterface = "InternalInterface"
)

var (
	/*
		Parsers can panic.
	*/

	NoParserRequired = func(v any) any { return v }
)

var (
	/*
		Note that cannot use parsers to parse/take values.
	*/

	DefaultTextValidator = func(v Value, d ValueDescriptor) error {
		if v.Type != TypeText {
			return fmt.Errorf("argument (%s) requires a value of type (%s), but a value of type (%s) entered", d.Name, d.Type, v.Type)
		}

		if v.Value == nil {
			return fmt.Errorf("argument (%s) is required", d.Name)
		}

		s, ok := v.Value.(string)
		if !ok {
			return fmt.Errorf("value of arguemnt (%s) is not a valid text", d.Name)
		}

		if d.Required && s == "" {
			return fmt.Errorf("argument (%s) is required", d.Name)
		}

		return nil
	}

	DefaultNumericValidator = func(v Value, d ValueDescriptor) error {
		if v.Type != TypeNumeric {
			return fmt.Errorf("argument (%s) requires a value of type (%s), but a value of type (%s) entered", d.Name, d.Type, v.Type)
		}

		if v.Value == nil {
			return fmt.Errorf("argument (%s) is required", d.Name)
		}

		return nil
	}

	DefaultValueValidator = func(v Value, d ValueDescriptor) error {
		switch d.Type {
		case TypeText:
			return DefaultTextValidator(v, d)
		case TypeNumeric:
			return DefaultNumericValidator(v, d)
		default:
			return nil
		}
	}
)

type ValueDescriptor struct {
	Name           string                             `json:"name"`
	Type           string                             `json:"type"`
	Required       bool                               `json:"required"`
	Default        any                                `json:"default,omitempty"`
	ValueValidator func(Value, ValueDescriptor) error `json:"-"`
}

func (v ValueDescriptor) Comparable() any {
	return v.Name
}

type Value struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Value       any           `json:"value,omitempty"`
	ValueParser func(any) any `json:"-"`
}

// ValueFromMap can panic.
func ValueFromMap(m map[string]any) Value {
	return Value{
		Name:  m["name"].(string),
		Type:  m["type"].(string),
		Value: m["value"].(any),
	}
}

func (v Value) Get() any {
	return v.ValueParser(v.Value)
}

func (v Value) Comparable() any {
	return v.Name
}
