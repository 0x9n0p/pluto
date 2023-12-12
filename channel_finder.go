package pluto

const ProcessorName_ChannelFinder = "CHANNEL_FINDER"

func init() {
	PredefinedProcessors[ProcessorName_ChannelFinder] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_ChannelFinder, &err)()
		return ChannelFinder{
			Name: Find("name", args...).Value.(string),
		}, err
	}
}

type ChannelFinder struct {
	Name string
}

func (p ChannelFinder) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_ChannelFinder},
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

func (p ChannelFinder) GetDescriptor() ProcessorDescriptor {
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
