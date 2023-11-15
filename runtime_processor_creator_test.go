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
					PredefinedProcessorName: "WRITE_TO_IO",
					AppendName:              "processor",
				},
			}},
		},
	})

	pluto.Process(&pluto.OutComingProcessable{
		Producer: TestIdentifier{
			Name: "TEST_PRODUCER",
			Kind: pluto.KindPipeline,
		},
		Consumer: TestIdentifier{
			Name: "CREATE_WRITE_TO_IO_PROCESSOR",
			Kind: pluto.KindPipeline,
		},
		Body: map[string]any{"io_interface": os.Stdout},
	})
}
