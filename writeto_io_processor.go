package pluto

import (
	"io"

	"go.uber.org/zap"
)

func init() {
	PredefinedProcessors["WRITE_TO_IO"] = func(r any) Processor {
		a, ok := r.(map[string]any)
		if !ok {
			ApplicationLogger.Debug(ApplicationLog{
				Message: "Cannot create WRITE_TO_IO processor: Body is not map[string]any",
			})
			return nil
		}

		v, found := a["io_interface"]
		if !found {
			ApplicationLogger.Debug(ApplicationLog{
				Message: "Cannot create WRITE_TO_IO processor: No io_interface found in body",
			})
			return nil
		}

		i, ok := v.(io.Writer)
		if !ok {
			ApplicationLogger.Debug(ApplicationLog{
				Message: "Cannot create WRITE_TO_IO processor: The io_interface is not writer",
			})
			return nil
		}

		return WriteToIOProcessor{
			Writer: i,
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
		Name:        "WRITE_TO_IO",
		Description: "",
		Input:       "",
		Output:      "",
	}
}
