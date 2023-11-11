package pluto

type ChannelJoinableProcessor struct {
	Name string
	Kind string
	Processor
}

func (c *ChannelJoinableProcessor) UniqueProperty() string {
	return c.Name
}

func (c *ChannelJoinableProcessor) PredefinedKind() string {
	return c.Kind
}
