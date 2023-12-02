package pluto

func init() {
	// TODO: Runtime processor creators should be added by HTTP APIs.
	PredefinedProcessors["RUNTIME_PROCESSOR_CREATOR_WRITE_TO_IO"] = func([]Value) Processor {
		return RuntimeProcessorCreator{
			PredefinedProcessorName: ProcessorName_WriteToInputOutput,
			AppendName:              "processor",
		}
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

	a[p.AppendName] = creator(processable.GetBody().([]Value))
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
