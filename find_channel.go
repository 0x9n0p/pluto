package pluto

const ProcessorName_FindChannel = "Find Channel"

func init() {
	PredefinedProcessors[ProcessorName_FindChannel] = func(args []Value) Processor {
		defer func() {
			ApplicationLogger.Error(ApplicationLog{
				Message: "Make sure you entered correct arguments",
				Extra:   map[string]any{"details": recover()},
			})
		}()

		return FindChannel{
			Name: Find("Name", args...).Value.(string),
		}
	}
}

type FindChannel struct {
	Name string
}

func (p FindChannel) Process(processable Processable) (Processable, bool) {
	for _, channel := range Channels {
		if channel.Name == p.Name {
			processable.SetBody(channel)
			return processable, true
		}
	}
	return processable, false
}

func (p FindChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{}
}

//func findChannel(name string) (Channel, bool) {
//	for _, channel := range Channels {
//		if channel.Name == name {
//			return channel, true
//		}
//	}
//	return Channel{}, false
//}
