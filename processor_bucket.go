package pluto

type ProcessorBucket struct {
	Processors []Processor
}

func (b *ProcessorBucket) Process(processable Processable) (new Processable, success bool) {
	for _, processor := range b.Processors {
		new, success = processor.Process(processable)
		if !success {
			return
		}
	}
	return
}

func (b *ProcessorBucket) Attach(processor Processor) {
	b.Processors = append(b.Processors, processor)
}

// TODO: DeAttach processors by processor.GetName
