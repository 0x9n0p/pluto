package pluto_test

import (
	"os"
	"pluto"
	"testing"
)

func TestRuntimeProcessorCreator(t *testing.T) {
	pluto.ReloadExecutionCache(map[string]pluto.Pipeline{
		"CREATE_WRITE_TO_IO_PROCESSOR": {
			Name: "CREATE_WRITE_TO_IO_PROCESSOR",
			ProcessorBucket: pluto.ProcessorBucket{Processors: []pluto.Processor{
				pluto.RuntimeProcessorCreator{
					PredefinedProcessorName: pluto.ProcessorName_WriteToInputOutput,
					AppendName:              "processor",
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
