package pluto

type FindChannel struct {
}

func (p FindChannel) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		return processable, false
	}

	v, found := appendable["channel_name"]
	if !found {
		return processable, false
	}

	channelName, ok := v.(string)
	if !ok {
		return processable, false
	}

	channel, found := findChannel(channelName)
	if !found {
		return processable, false
	}

	appendable["channel"] = channel
	processable.SetBody(appendable)

	return processable, true
}

func (p FindChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "FIND_CHANNEL",
		Description: "",
		Input:       "channel_name",
		Output:      "channel",
	}
}

func findChannel(name string) (Channel, bool) {
	for _, channel := range Channels {
		if channel.Name == name {
			return channel, true
		}
	}
	return Channel{}, false
}
