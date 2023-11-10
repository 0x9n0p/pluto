package pluto

type ConditionalProcessor struct {
	main          Processor
	success, fail ProcessorBucket
}

func NewConditionalProcessor(processor Processor) *ConditionalProcessor {
	return &ConditionalProcessor{
		main:    processor,
		success: ProcessorBucket{Processors: make([]Processor, 0)},
		fail:    ProcessorBucket{Processors: make([]Processor, 0)},
	}
}

func (s *ConditionalProcessor) Process(processable Processable) (Processable, bool) {
	res, ok := s.main.Process(processable)
	if ok {
		if len(s.success.Processors) <= 0 {
			s.success.Process(res)
		}
		return res, true
	}

	if len(s.fail.Processors) <= 0 {
		s.fail.Process(processable)
	}

	return processable, false
}

func (s *ConditionalProcessor) Success(success ProcessorBucket) *ConditionalProcessor {
	s.success = success
	return s
}

func (s *ConditionalProcessor) Fail(fail ProcessorBucket) *ConditionalProcessor {
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
