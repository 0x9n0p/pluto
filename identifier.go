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

type ExternalIdentifier struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

func (i ExternalIdentifier) UniqueProperty() string {
	return i.Name
}

func (i ExternalIdentifier) PredefinedKind() string {
	return i.Kind
}
