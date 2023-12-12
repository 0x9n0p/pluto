package pluto

//
//const ProcessorName_ExecAndJoinResult = "Execute processor and join the result"
//
//func init() {
//	PredefinedProcessors[ProcessorName_ExecAndJoinResult] = func(args []Value) (p Processor, err error) {
//		defer creatorPanicHandler(ProcessorName_ExecAndJoinResult, &err)()
//		return ExecAndJoinResult{
//			Processor: Find("Processor", args...).Get().(Processor),
//		}, err
//	}
//}
//
//type ExecAndJoinResult struct {
//	Processor Processor
//}
//
//func (p ExecAndJoinResult) Process(processable Processable) (Processable, bool) {
//	a, ok := processable.GetBody().(map[string]any)
//	if !ok {
//		ApplicationLogger.Debug(ApplicationLog{
//			Message: "The body does not support the append operation",
//		})
//		return processable, false
//	}
//
//	r, ok := p.Processor.Process(processable)
//	if !ok {
//		return r, false
//	}
//
//	{
//		ar, ok := r.GetBody().(map[string]any)
//		if !ok {
//			ApplicationLogger.Debug(ApplicationLog{
//				Message: "The result body does not support the append operation",
//			})
//			return r, false
//		}
//
//		for k, v := range a {
//			ar[k] = v
//		}
//
//		r.SetBody(ar)
//	}
//
//	return r, true
//}
//
//func (p ExecAndJoinResult) GetDescriptor() ProcessorDescriptor {
//	return ProcessorDescriptor{}
//}
