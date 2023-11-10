package pluto

type Pipeline struct {
	Name string
	ProcessorBucket
}

func (p *Pipeline) GetName() string {
	return p.Name
}
