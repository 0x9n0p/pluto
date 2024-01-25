package pluto

const ProcessorName_Fork = "FORK"

func init() {
	PredefinedProcessors[ProcessorName_Fork] = func(args []Value) (p Processor, err error) {
		defer CreatorPanicHandler(ProcessorName_Fork, &err)()
		return Fork{
			Name: Find("name", args...).Get().(string),
		}, err
	}
}

type Fork struct {
	Name string
}

func (p Fork) Process(processable Processable) (Processable, bool) {
	go (&Execute{Name: p.Name}).Process(processable)
	return processable, true
}
