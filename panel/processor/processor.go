package processor

type Descriptor struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Icon        string                     `json:"icon"`
	Arguments   map[string]ValueDescriptor `json:"arguments"`
	Input       map[string]ValueDescriptor `json:"input"`
	Output      map[string]ValueDescriptor `json:"output"`
	Category    string                     `json:"category"`
}

type Processor struct {
	Name      string           `json:"name"`
	Arguments map[string]Value `json:"arguments"`
}
