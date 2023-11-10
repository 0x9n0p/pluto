package pluto_test

import (
	"fmt"
	"pluto"
)

type PrintProcessor struct {
}

func (p PrintProcessor) Process(processable pluto.Processable) (pluto.Processable, bool) {
	fmt.Printf("%s: %s\n", p.GetDescriptor().Name, processable)
	return processable, true
}

func (p PrintProcessor) GetDescriptor() pluto.ProcessorDescriptor {
	return pluto.ProcessorDescriptor{
		Name:        "Print Processor",
		Description: "Description",
		Input:       "",
		Output:      "",
	}
}
