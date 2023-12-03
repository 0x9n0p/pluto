package pluto

const ProcessorName_CreateChannel = "Create Channel"

func init() {
	PredefinedProcessors[ProcessorName_CreateChannel] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_ExecAndJoinResult, &err)
		return CreateChannel{
			Name:   Find("Name", args...).Value.(string),
			Length: Find("Length", args...).Value.(uint),
		}, err
	}
}

type CreateChannel struct {
	Name   string
	Length uint
}

func (p CreateChannel) Process(processable Processable) (Processable, bool) {
	channel := NewChannel(p.Name, p.Length)
	Channels[channel.ID] = channel
	processable.SetBody(channel)
	return processable, true
}

func (p CreateChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{}
}
