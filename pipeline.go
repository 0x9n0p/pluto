package pluto

import (
	"net/http"
)

var (
	PipelineNotFound = &Error{
		Code:     1,
		HTTPCode: http.StatusNotFound,
		Message:  "Pipeline not found",
	}
)

type Pipeline struct {
	Name            string `json:"name"`
	ProcessorBucket `json:"processor_bucket"`
}

type PipelineIdentifier struct {
	Name string `json:"name"`
}

func (p *PipelineIdentifier) GetPipeline() (Pipeline, error) {
	return Pipeline{}, PipelineNotFound
}

func (p *PipelineIdentifier) UniqueProperty() string {
	return p.Name
}

func (p *PipelineIdentifier) PredefinedKind() string {
	return KindPipeline
}
