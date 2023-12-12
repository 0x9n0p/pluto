package pluto

import "fmt"

const ProcessorName_ChannelCreator = "CHANNEL_CREATOR"

func init() {
	PredefinedProcessors[ProcessorName_ChannelCreator] = func(args []Value) (p Processor, err error) {
		defer func() {
			if v := recover(); v != nil {
				err = fmt.Errorf("make sure you enter the arguments of (%s) correctly: %s", ProcessorName_ChannelCreator, v.(error))
			}
		}()
		return ChannelCreator{
			Name:   Find("name", args...).Get().(string),
			Length: Find("length", args...).Get().(int),
		}, err
	}
}

type ChannelCreator struct {
	Name   string
	Length int
}

func (p ChannelCreator) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_ChannelCreator},
		})
		return processable, false
	}

	channel := NewChannel(p.Name, p.Length)
	Channels[channel.ID] = channel
	appendable["channel"] = channel
	processable.SetBody(appendable)
	return processable, true
}

func (p ChannelCreator) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{}
}
