package pluto

import "fmt"

/*
	RUNTIME_PROCESSOR_CREATOR can not pass arguments to the processor creator.
*/

const ProcessorName_RuntimeProcessorCreator = "RUNTIME_PROCESSOR_CREATOR"

func init() {
	PredefinedProcessors[ProcessorName_RuntimeProcessorCreator] = func(args []Value) (p Processor, err error) {
		defer CreatorPanicHandler(ProcessorName_RuntimeProcessorCreator, &err)()
		return RuntimeProcessorCreator{
			ProcessorName: Find("processor_name", args...).Get().(string),
			AppendName:    Find("append_name", args...).Get().(string),
		}, err
	}
}

type RuntimeProcessorCreator struct {
	ProcessorName string
	AppendName    string
}

func (p RuntimeProcessorCreator) Process(processable Processable) (Processable, bool) {
	a, ok := processable.GetBody().(map[string]any)
	if !ok {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "The body does not support the append operation",
			Extra:   map[string]any{"issuer": ProcessorName_RuntimeProcessorCreator},
		})
		return processable, false
	}

	creator, found := PredefinedProcessors[p.ProcessorName]
	if !found {
		return processable, false
	}

	// TODO: Should we take/parse the arguments from the body?
	processor, err := creator([]Value{})
	if err != nil {
		ApplicationLogger.Error(ApplicationLog{
			Message: fmt.Sprintf("Runtime processor creator failed to create the processor (%s)", p.ProcessorName),
			Extra:   map[string]any{"details": err},
		})
		return processable, false
	}

	a[p.AppendName] = processor
	processable.SetBody(a)

	return processable, true
}
