package pluto

func init() {
	// TODO: Runtime processor creators should be added by HTTP APIs.
	PredefinedProcessors["RUNTIME_PROCESSOR_CREATOR_WRITE_TO_IO"] = func(any) Processor {
		return RuntimeProcessorCreator{
			PredefinedProcessorName: "WRITE_TO_IO",
			AppendName:              "processor",
		}
	}
}

type RuntimeProcessorCreator struct {
	PredefinedProcessorName string
	AppendName              string
}

func (p RuntimeProcessorCreator) Process(processable Processable) (Processable, bool) {
	a, ok := processable.GetBody().(Appendable)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Body is not appendable",
		})
		return processable, false
	}

	creator, found := PredefinedProcessors[p.PredefinedProcessorName]
	if !found {
		return processable, false
	}

	a[p.AppendName] = creator(processable.GetBody())
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
