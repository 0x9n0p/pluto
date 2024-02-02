package pipeline

import (
	"encoding/json"
	"net/http"
	"pluto"
	"pluto/panel/database"
	"pluto/panel/processor"
	"time"

	"go.uber.org/zap"
)

type Pipeline struct {
	Name             string                `json:"name" validate:"required"`
	ExecuteOnStartup bool                  `json:"execute_on_startup"  validate:"boolean"`
	Processors       []processor.Processor `json:"processors" validate:"required"`
	SavedAt          time.Time             `json:"saved_at"`

	Transaction *database.Transaction `json:"-"`
}

func (p *Pipeline) Create() (pluto.Pipeline, error) {
	pipeline := pluto.Pipeline{
		Name:             p.Name,
		ExecuteOnStartup: p.ExecuteOnStartup,
		ProcessorBucket:  pluto.ProcessorBucket{Processors: []pluto.Processor{}},
	}

	for _, wantedProcessor := range p.Processors {
		createdProcessor, err := wantedProcessor.Create()
		if err != nil {
			return pluto.Pipeline{}, err
		}

		pipeline.Processors = append(pipeline.Processors, createdProcessor)
	}

	return pipeline, nil
}

func (p *Pipeline) Active() error {
	// TODO
	panic("Implement me")
}

func (p *Pipeline) Delete() error {
	err := p.Transaction.Bucket(bucket).Delete([]byte(p.Name))
	if err != nil {
		pluto.Log.Error("Failed to delete pipeline", zap.String("pipeline_name", p.Name), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to delete pipeline",
		}
	}
	return nil
}

func (p *Pipeline) Save() error {
	p.SavedAt = time.Now()

	// Validation :)
	if _, err := p.Create(); err != nil {
		return err
	}

	b, err := json.Marshal(p)
	if err != nil {
		pluto.Log.Error("Failed to marshal the pipeline", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save pipeline",
		}
	}

	err = p.Transaction.Bucket(bucket).Put([]byte(p.Name), b)
	if err != nil {
		pluto.Log.Error("Failed to put pipeline", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to save pipeline",
		}
	}

	return nil
}
