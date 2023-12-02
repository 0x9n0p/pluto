package pluto

const ProcessorName_CreateChannel = "Create Channel"

func init() {
	PredefinedProcessors[ProcessorName_CreateChannel] = func(args []Value) Processor {
		defer func() {
			ApplicationLogger.Error(ApplicationLog{
				Message: "Make sure you entered correct arguments",
				Extra:   map[string]any{"details": recover()},
			})
		}()

		return CreateChannel{
			Name:   Find("Name", args...).Value.(string),
			Length: Find("Length", args...).Value.(uint),
		}
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
