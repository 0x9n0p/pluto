package pluto

const ProcessorName_JoinChannel = "Join Channel"

func init() {
	PredefinedProcessors[ProcessorName_JoinChannel] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_JoinChannel, &err)()
		return JoinChannel{}, err
	}
}

type JoinChannel struct {
}

func (p JoinChannel) Process(processable Processable) (Processable, bool) {
	appendable, ok := processable.GetBody().(map[string]any)
	if !ok {
		return processable, false
	}

	v, found := appendable["channel"]
	if !found {
		return processable, false
	}

	channel, ok := v.(Channel)
	if !ok {
		return processable, false
	}

	v, found = appendable["identifier"]
	if !found {
		return processable, false
	}

	identifier, ok := v.(Identifier)
	if !ok {
		return processable, false
	}

	v, found = appendable["processor"]
	if !found {
		return processable, false
	}

	processor, ok := v.(Processor)
	if !ok {
		return processable, false
	}

	channel.Join(&BaseJoinable{identifier, processor})

	return processable, true
}

func (p JoinChannel) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{}
}

type BaseJoinable struct {
	Identifier
	Processor
}
