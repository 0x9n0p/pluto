package pluto

type RuntimeProcessorCreator struct {
	PredefinedProcessorName string
	AppendName              string
}

func (p RuntimeProcessorCreator) Process(processable Processable) (Processable, bool) {
	a, ok := processable.GetBody().(Appendable)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Body is not appendable",
			Extra:   map[string]any{"producer": processable.GetProducer()},
		})
		return processable, false
	}

	creator, found := PredefinedProcessors[p.PredefinedProcessorName]
	if !found {
		return processable, false
	}

	a[p.AppendName] = creator(processable)
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
