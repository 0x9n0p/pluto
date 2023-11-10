package pluto_test

import (
	"pluto"
	"testing"
)

func TestConditionalProcessor(t *testing.T) {
	p := pluto.Pipeline{
		Name: "PRINT_PIPELINE",
		ProcessorBucket: pluto.ProcessorBucket{
			Processors: []pluto.Processor{
				pluto.NewConditionalProcessor(&PrintProcessor{}).
					Success(
						pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
							return processable, true
						}),
					),
			},
		},
	}

	if _, success := p.Process([]byte("Hello World")); !success {
		t.FailNow()
	}
}
