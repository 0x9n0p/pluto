package pluto

type Config struct {
	Pipelines []struct {
		Name       string `json:"name"`
		Processors []struct {
			Name    string   `json:"name"`
			Success []string `json:"success,omitempty"`
			Fail    []string `json:"fail,omitempty"`
		} `json:"processors"`
	} `json:"pipelines"`
}

// ResolveConfig REFACTOR
func ResolveConfig(config Config) (o map[string]Pipeline) {
	for _, pipeline := range config.Pipelines {
		p := Pipeline{
			Name:            pipeline.Name,
			ProcessorBucket: ProcessorBucket{},
		}

		for _, processor := range pipeline.Processors {
			processorCreator, found := PredefinedProcessors[processor.Name]
			if !found {
				ApplicationLogger.Warning(ApplicationLog{
					Message: "Predefined processor not found to attach to pipeline",
					Extra:   map[string]any{"processor_name": processor.Name},
				})
				continue
			}

			var conditionalProcessor ConditionalProcessor
			conditionalProcessor.main = processorCreator()

			if processor.Success == nil || len(processor.Success) <= 0 {
				conditionalProcessor.success = ProcessorBucket{make([]Processor, 0)}

				for _, successProcessor := range processor.Success {
					processorCreator, found := PredefinedProcessors[successProcessor]
					if !found {
						ApplicationLogger.Warning(ApplicationLog{
							Message: "Predefined processor not found to attach to success path",
							Extra:   map[string]any{"processor_name": successProcessor},
						})
						continue
					}

					conditionalProcessor.success.Attach(processorCreator())
				}
			}

			if processor.Fail == nil || len(processor.Fail) <= 0 {
				conditionalProcessor.fail = ProcessorBucket{make([]Processor, 0)}

				for _, failProcessor := range processor.Fail {
					processorCreator, found := PredefinedProcessors[failProcessor]
					if !found {
						ApplicationLogger.Warning(ApplicationLog{
							Message: "Predefined processor not found to attach to fail path",
							Extra:   map[string]any{"processor_name": failProcessor},
						})
						continue
					}

					conditionalProcessor.fail.Attach(processorCreator())
				}
			}

			p.ProcessorBucket.Attach(&conditionalProcessor)
		}

		o[pipeline.Name] = p
	}

	return
}
