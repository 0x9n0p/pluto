package pluto

const ProcessorName_Execute = "EXECUTE"

func init() {
	PredefinedProcessors[ProcessorName_Execute] = func(args []Value) (p Processor, err error) {
		defer CreatorPanicHandler(ProcessorName_Execute, &err)()
		return Execute{
			Name:         Find("name", args...).Get().(string),
			AppendResult: Find("append_result", args...).Get().(bool),
		}, err
	}
}

type Execute struct {
	Name         string
	AppendResult bool
}

func (p Execute) Process(processable Processable) (Processable, bool) {
	ExecutionCacheMutex.RLock()

	pipeline, found := ExecutionCache[p.Name]
	if !found {
		ExecutionCacheMutex.RUnlock()
		ApplicationLogger.Warning(ApplicationLog{
			Message: "Pipeline not found",
			Extra:   map[string]any{"unique_property": p.Name},
		})
		return processable, false
	}

	ExecutionCacheMutex.RUnlock()

	{
		result, success := pipeline.Process(processable)

		if p.AppendResult {
			a, ok := processable.GetBody().(map[string]any)
			if !ok {
				ApplicationLogger.Debug(ApplicationLog{
					Message: "The body does not support the append operation",
					Extra:   map[string]any{"issuer": ProcessorName_Execute},
				})
				return processable, false
			}

			r, ok := result.GetBody().(map[string]any)
			if !ok {
				ApplicationLogger.Debug(ApplicationLog{
					Message: "The result body does not support the append operation",
					Extra:   map[string]any{"issuer": ProcessorName_Execute},
				})
				return processable, false
			}

			for k, v := range a {
				r[k] = v
			}

			result.SetBody(r)
		}

		return result, success
	}
}
