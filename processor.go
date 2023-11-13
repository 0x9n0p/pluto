package pluto

var PredefinedProcessors = make(map[string]func(any) Processor)

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

type EmptyProcessor struct {
}

func (p EmptyProcessor) Process(processable Processable) (Processable, bool) {
	return processable, true
}

func (p EmptyProcessor) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name: "Empty Processor",
	}
}
