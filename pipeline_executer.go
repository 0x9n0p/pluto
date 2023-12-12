package pluto

const ProcessorName_PipelineExecuter = "PIPELINE_EXECUTER"

func init() {
	PredefinedProcessors[ProcessorName_PipelineExecuter] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_PipelineExecuter, &err)()
		return PipelineExecuter{
			Name:         Find("name", args...).Get().(string),
			AppendResult: Find("append_result", args...).Get().(bool),
		}, err
	}
}

type PipelineExecuter struct {
	Name         string
	AppendResult bool
}

func (p PipelineExecuter) Process(processable Processable) (Processable, bool) {
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
					Extra:   map[string]any{"issuer": ProcessorName_PipelineExecuter},
				})
				return processable, false
			}

			r, ok := result.GetBody().(map[string]any)
			if !ok {
				ApplicationLogger.Debug(ApplicationLog{
					Message: "The result body does not support the append operation",
					Extra:   map[string]any{"issuer": ProcessorName_PipelineExecuter},
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
