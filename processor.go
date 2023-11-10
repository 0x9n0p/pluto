package pluto

type Processor interface {
	GetDescriptor() ProcessorDescriptor
	// Process
	// The boolean indicates that the next processor can be executed or not.
	Process(Processable) (Processable, bool)
}

type ProcessorDescriptor struct {
	Name        string
	Description string
	Input       string
	Output      string
}
