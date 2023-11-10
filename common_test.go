package pluto_test

import (
	"fmt"
	"pluto"
)

type PrintProcessor struct {
}

func (p PrintProcessor) GetName() string {
	return "PRINT_PROCESSOR"
}

func (p PrintProcessor) Process(processable pluto.Processable) (pluto.Processable, bool) {
	fmt.Printf("%s: %s\n", p.GetName(), processable)
	return processable, true
}
