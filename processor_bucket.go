package pluto

type ProcessorBucket struct {
	Processors []Processor
}

func (b *ProcessorBucket) Attach(processor Processor) {
	b.Processors = append(b.Processors, processor)
}
