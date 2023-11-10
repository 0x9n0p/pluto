package pluto

type Config struct {
	Pipelines []struct {
		Name       string   `json:"name"`
		Processors []string `json:"processors"`
	} `json:"pipelines"`
}

func ResolveConfig(config Config) (o map[string]Pipeline) {
	for _, pipeline := range config.Pipelines {
		p := Pipeline{
			Name:            pipeline.Name,
			ProcessorBucket: ProcessorBucket{},
		}

		for _, processorName := range pipeline.Processors {
			processorCreator, found := PredefinedProcessors[processorName]
			if !found {
				ApplicationLogger.Warning(ApplicationLog{
					Message: "Predefined processor not found to attach to pipeline",
					Extra:   map[string]any{"processor_name": processorName},
				})
				continue
			}

			p.ProcessorBucket.Attach(processorCreator())
		}

		o[pipeline.Name] = p
	}

	return
}
