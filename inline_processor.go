package pluto

type InlineProcessor struct {
	f func(processable Processable) (Processable, bool)
}

func NewInlineProcessor(f func(processable Processable) (Processable, bool)) *InlineProcessor {
	return &InlineProcessor{
		f: f,
	}
}

func (s *InlineProcessor) Process(processable Processable) (Processable, bool) {
	return s.f(processable)
}
