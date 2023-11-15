package pluto

type AppendResultProcessor struct {
	Processor Processor
}

func (p AppendResultProcessor) Process(processable Processable) (Processable, bool) {
	a, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Body is not map[string]any",
		})
		return processable, false
	}

	r, ok := p.Processor.Process(processable)
	if !ok {
		return r, false
	}

	{
		ar, ok := r.GetBody().(map[string]any)
		if !ok {
			ApplicationLogger.Debug(ApplicationLog{
				Message: "Body of result is not map[string]any",
			})
			return r, false
		}

		for k, v := range a {
			ar[k] = v
		}

		r.SetBody(ar)
	}

	return r, true
}

func (p AppendResultProcessor) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "APPEND_RESULT",
		Description: "",
		Input:       "",
		Output:      "",
	}
}
