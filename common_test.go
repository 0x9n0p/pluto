package pluto_test

func init() {
	//pluto.PredefinedProcessors["Print Processor"] = func([]pluto.Value) pluto.Processor { return &PrintProcessor{} }
}

//
//type PrintProcessor struct {
//}
//
//func (p PrintProcessor) Process(processable pluto.Processable) (pluto.Processable, bool) {
//	fmt.Printf("%s: %s\n", p.GetDescriptor().Name, processable)
//	return processable, true
//}
//
//type TestIdentifier struct {
//	Name string
//	Kind string
//}
//
//func (i TestIdentifier) UniqueProperty() string {
//	return i.Name
//}
//
//func (i TestIdentifier) PredefinedKind() string {
//	return i.Kind
//}
