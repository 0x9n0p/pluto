package pluto

const ProcessorName_FindChannel = "Find Channel"

func init() {
	PredefinedProcessors[ProcessorName_FindChannel] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_FindChannel, &err)()
		return FindChannel{
			Name: Find("Name", args...).Value.(string),
		}, err
	}
}

type FindChannel struct {
	Name string
}

func (p FindChannel) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_FindChannel},
		})
		return processable, false
	}

	for _, channel := range Channels {
		if channel.Name == p.Name {
			appendable["channel"] = channel
			processable.SetBody(appendable)
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
