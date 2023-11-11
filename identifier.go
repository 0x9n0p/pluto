package pluto

const (
	KindChannel             = "Channel"
	KindPipeline            = "Pipeline"
	KindInternalProcessable = "InternalProcessable"
)

type Identifier interface {
	UniqueProperty() string
	PredefinedKind() string
}

func CompareIdentifiers(first Identifier, second Identifier) bool {
	return first.PredefinedKind() == second.PredefinedKind() &&
		first.UniqueProperty() == second.UniqueProperty()
}
