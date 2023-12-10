package pluto

import "fmt"

const ProcessorName_CreateChannel = "Create Channel"

func init() {
	PredefinedProcessors[ProcessorName_CreateChannel] = func(args []Value) (p Processor, err error) {
		defer func() {
			if v := recover(); v != nil {
				err = fmt.Errorf("make sure you enter the arguments of (%s) correctly: %s", ProcessorName_CreateChannel, v.(error))
			}
		}()
		return CreateChannel{
			Name:   Find("Name", args...).Get().(string),
			Length: Find("Length", args...).Get().(int),
		}, err
	}
}

type CreateChannel struct {
	Name   string
	Length int
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
