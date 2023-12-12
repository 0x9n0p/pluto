package pluto

import (
	"io"

	"go.uber.org/zap"
)

const ProcessorName_IOWriter = "IO_WRITER"

func init() {
	PredefinedProcessors[ProcessorName_IOWriter] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_IOWriter, &err)()
		return IOWriter{
			Writer: Find("io_interface", args...).Value.(io.Writer),
		}, err
	}
}

type IOWriter struct {
	io.Writer
}

func (p IOWriter) Process(processable Processable) (Processable, bool) {
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

func (p IOWriter) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        ProcessorName_IOWriter,
		Description: "",
		Input:       "",
		Output:      "",
	}
}