package pluto

type ConditionalProcessor struct {
	processor     Processor
	success, fail Processor

	main Pipeline
}

func NewConditionalProcessor(processor Processor) *ConditionalProcessor {
	return &ConditionalProcessor{
		processor: processor,
	}
}

func (s *ConditionalProcessor) Process(processable Processable) (Processable, bool) {
	res, ok := s.processor.Process(processable)
	if ok {
		if s.success != nil {
			s.success.Process(res)
		}
		return res, true
	}

	if s.fail != nil {
		s.fail.Process(processable)
	}

	return processable, false
}

func (s *ConditionalProcessor) Success(success Processor) *ConditionalProcessor {
	s.success = success
	return s
}

func (s *ConditionalProcessor) Fail(fail Processor) *ConditionalProcessor {
	s.fail = fail
	return s
}

func (s *ConditionalProcessor) GetDescriptor() ProcessorDescriptor {
	return ProcessorDescriptor{
		Name:        "Conditional Processor",
		Description: "Description",
		Input:       "",
		Output:      "",
	}
}
