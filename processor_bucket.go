package pluto

type ProcessorBucket struct {
	Processors []Processor
}

func (b *ProcessorBucket) Process(processable Processable) (Processable, bool) {
	for _, processor := range b.Processors {
		p, success := processor.Process(processable)
		processable = p

		if !success {
			return processable, false
		}
	}
	return processable, true
}

func (b *ProcessorBucket) Attach(processor Processor) {
	b.Processors = append(b.Processors, processor)
}

// TODO: DeAttach processors by processor.GetName
