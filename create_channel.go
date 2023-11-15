package pluto

type CreateChannel struct {
}

func (p CreateChannel) Process(processable Processable) (Processable, bool) {
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

	v, found = appendable["channel_length"]
	if !found {
		return processable, false
	}

	channelLength, ok := v.(uint)
	if !ok {
		return processable, false
	}

	channel := NewChannel(channelName, channelLength)
	Channels[channel.ID] = channel

	appendable["channel"] = channel
	processable.SetBody(appendable)

	return processable, true
}

func (p CreateChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "CREATE_CHANNEL",
		Description: "",
		Input:       "channel_name,channel_length",
		Output:      "channel",
	}
}
