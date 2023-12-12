package pluto

const ProcessorName_PipelineExecuter = "PIPELINE_EXECUTER"

func init() {
	PredefinedProcessors[ProcessorName_PipelineExecuter] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_PipelineExecuter, &err)()
		return PipelineExecuter{
			Name: Find("name", args...).Value.(string),
		}, err
	}
}

type PipelineExecuter struct {
	Name string
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

	return pipeline.Process(processable)
}
