package pipeline

import (
	"fmt"
	"pluto"
	"pluto/panel/processor"
	"time"

	"go.uber.org/zap"
)

type Pipeline struct {
	Name       string                `json:"name" validate:"required"`
	Processors []processor.Processor `json:"processors" validate:"required"`
	SavedAt    time.Time             `json:"saved_at"`
}

func (p *Pipeline) Create() (pluto.Pipeline, error) {
	pipeline := pluto.Pipeline{
		Name:            p.Name,
		ProcessorBucket: pluto.ProcessorBucket{Processors: []pluto.Processor{}},
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
	if err := GetStorage().delete(p.Name); err != nil {
		pluto.Log.Error("Delete pipeline", zap.String("pipeline_name", p.Name), zap.Error(err))
		return fmt.Errorf("storage: %v", err)
	}
	return nil
}

func (p *Pipeline) Save() error {
	// Validation :)
	if _, err := p.Create(); err != nil {
		return err
	}

	if err := GetStorage().save(p); err != nil {
		pluto.Log.Error("Save pipeline", zap.String("pipeline_name", p.Name), zap.Error(err))
		return fmt.Errorf("storage: %v", err)
	}

	return nil
}
