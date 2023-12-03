package pluto

import "fmt"

func init() {
	// TODO: Runtime processor creators should be added by HTTP APIs.
	PredefinedProcessors["RUNTIME_PROCESSOR_CREATOR_WRITE_TO_IO"] = func([]Value) (p Processor, err error) {
		defer creatorPanicHandler("RUNTIME_PROCESSOR_CREATOR_WRITE_TO_IO", &err)()
		return RuntimeProcessorCreator{
			PredefinedProcessorName: ProcessorName_WriteToInputOutput,
			AppendName:              "processor",
		}, err
	}
}

// RuntimeProcessorCreator
// Deprecated
type RuntimeProcessorCreator struct {
	PredefinedProcessorName string
	AppendName              string
}

func (p RuntimeProcessorCreator) Process(processable Processable) (Processable, bool) {
	a, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Body is not map[string]any",
		})
		return processable, false
	}

	creator, found := PredefinedProcessors[p.PredefinedProcessorName]
	if !found {
		return processable, false
	}

	processor, err := creator(processable.GetBody().([]Value))
	if err != nil {
		ApplicationLogger.Error(ApplicationLog{
			Message: fmt.Sprintf("Runtime processor creator failed to create the processor (%s)", p.PredefinedProcessorName),
			Extra:   map[string]any{"details": err},
		})
		return processable, false
	}

	a[p.AppendName] = processor
	processable.SetBody(a)

	return processable, true
}

func (p RuntimeProcessorCreator) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "RUNTIME_PROCESSOR_CREATOR",
		Description: "",
		Input:       "",
		Output:      "",
	}
}
