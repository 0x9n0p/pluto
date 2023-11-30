package pluto

import (
	"io"

	"go.uber.org/zap"
)

const ProcessorName_WriteToInputOutput = "Write to input/output"

func init() {
	PredefinedProcessors[ProcessorName_WriteToInputOutput] = func(args []Value) Processor {
		defer func() {
			// TODO: Application log
		}()

		return WriteToIOProcessor{
			Writer: Find("io_interface", args...).Value.(io.Writer),
		}
	}
}

type WriteToIOProcessor struct {
	io.Writer
}

func (p WriteToIOProcessor) Process(processable Processable) (Processable, bool) {
	b, ok := processable.GetBody().([]byte)
	if !ok {
		Log.Error("Channels only support []byte to publish")
		return processable, false
	}

	_, err := p.Write(b)
	if err != nil {
		Log.Debug("Write to io", zap.Error(err))
		return processable, false
	}

	return processable, true
}

func (p WriteToIOProcessor) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        ProcessorName_WriteToInputOutput,
		Description: "",
		Input:       "",
		Output:      "",
	}
}
