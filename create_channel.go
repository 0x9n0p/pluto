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
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_CreateChannel},
		})
		return processable, false
	}

	channel := NewChannel(p.Name, p.Length)
	Channels[channel.ID] = channel
	appendable["channel"] = channel
	processable.SetBody(appendable)
	return processable, true
}

func (p CreateChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{}
}
