package pluto_test

import (
	"fmt"
	"pluto"
	"testing"
)

func TestPipeline(t *testing.T) {
	p := pluto.Pipeline{
		Name: "PRINT_PIPELINE",
		ProcessorBucket: pluto.ProcessorBucket{
			Processors: []pluto.Processor{
				PrintProcessor{},
			},
		},
	}

	if _, success := p.Process([]byte("Hello World")); !success {
		t.FailNow()
	}
}

type PrintProcessor struct {
}

func (p PrintProcessor) GetName() string {
	return "PRINT_PROCESSOR"
}

func (p PrintProcessor) Process(processable pluto.Processable) (pluto.Processable, bool) {
	fmt.Printf("%s: %s\n", p.GetName(), processable)
	return processable, true
}
