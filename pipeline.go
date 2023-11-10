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
