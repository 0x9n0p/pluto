package pluto

import "fmt"

var PredefinedProcessors = make(map[string]func([]Value) (Processor, error))

func CreatorPanicHandler(processorName string, err *error) func() {
	return func() {
		if v := recover(); v != nil {
			*err = fmt.Errorf("make sure you enter the arguments of (%s) correctly: %s", processorName, v.(error))
		}
	}
}

type Processor interface {
	// Process
	// The boolean indicates that the next processor can be executed or not.
	Process(Processable) (Processable, bool)
}

type EmptyProcessor struct {
}

func (p EmptyProcessor) Process(processable Processable) (Processable, bool) {
	return processable, true
}
