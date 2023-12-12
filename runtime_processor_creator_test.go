package pluto_test

import (
	"os"
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
		Producer: pluto.ExternalIdentifier(TestIdentifier{
			Name: "TEST_PRODUCER",
			Kind: pluto.KindPipeline,
		}),
		Consumer: pluto.ExternalIdentifier(TestIdentifier{
			Name: "CREATE_WRITE_TO_IO_PROCESSOR",
			Kind: pluto.KindPipeline,
		}),
		Body: []pluto.Value{
			{
				Name:  "io_interface",
				Type:  pluto.TypeInternalInterface,
				Value: os.Stdout,
			},
		},
	})
}
