package processor

var Processors = []Descriptor{
	{
		Name:        "Create Table",
		Description: "This is for test",
		Icon:        "https://...",
		Arguments: map[string]ValueDescriptor{
			"create_if_not_exists": {
				Type:    TypeNumeric,
				Default: 0,
			},
			"table_name": {
				Type:     TypeText,
				Required: true,
			},
			"columns": {
				Type: TypeList,
				//Default:  []Value{},
				Required: true,
			},
		},
		Input: map[string]ValueDescriptor{},
		Output: map[string]ValueDescriptor{
			"table_id": {
				Type: TypeText,
			},
		},
		Category: "Storage",
	},

	{
		Name:        "MY_PROCESSOR",
		Description: "This is for test",
		Icon:        "https://...",
		Arguments:   map[string]ValueDescriptor{},
		Input:       map[string]ValueDescriptor{},
		Output:      map[string]ValueDescriptor{},
		Category:    "Storage",
	},
}
