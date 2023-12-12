package pluto_test

import (
	"pluto"
	"testing"
)

func TestRuntimeProcessorCreator(t *testing.T) {
	pluto.ReloadExecutionCache(map[string]pluto.Pipeline{
		"TEST_PIPELINE": {
			Name: "TEST_PIPELINE",
			ProcessorBucket: pluto.ProcessorBucket{Processors: []pluto.Processor{
				pluto.RuntimeProcessorCreator{
					ProcessorName: pluto.ProcessorName_IOWriter,
					AppendName:    "processor",
				},
			}},
		},
	})

	pluto.Process(&pluto.OutComingProcessable{
		Producer: pluto.ExternalIdentifier{
			Name: "TEST_PRODUCER",
			Kind: pluto.KindPipeline,
		},
		Consumer: pluto.ExternalIdentifier{
			Name: "TEST_PIPELINE",
			Kind: pluto.KindPipeline,
		},
		Body: map[string]any{},
	})
}
