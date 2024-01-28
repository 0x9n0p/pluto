package tcp

import "pluto"

const ProcessorName_SendResponse = "SEND_RESPONSE"

func init() {
	pluto.PredefinedProcessors[ProcessorName_SendResponse] = func(args []pluto.Value) (p pluto.Processor, err error) {
		defer pluto.creatorPanicHandler(ProcessorName_SendResponse, &err)()
		return SendResponse{
			PipelineName: pluto.Find("pipeline_name", args...).Get().(string),
		}, err
	}
}

// SendResponse Encodes the body, sends it back to the producer and never changes the input processable.
type SendResponse struct {
	// TODO
	// Filter []string
	PipelineName string
}

func (p SendResponse) Process(processable pluto.Processable) (pluto.Processable, bool) {
	outGoing := pluto.OutGoingProcessable{
		Consumer: pluto.ExternalIdentifier{
			Name: p.PipelineName,
			Kind: pluto.KindPipeline,
		},
		Body: processable.GetBody(),
	}

	outComing, ok := processable.(*pluto.OutComingProcessable)
	if !ok {
		pluto.ApplicationLogger.Error(pluto.ApplicationLog{
			Message: "The processable is not an out_coming_processable",
			Extra:   map[string]any{"issuer": ProcessorName_SendResponse},
		})
		return processable, false
	}

	if err := outComing.Encoder.Encode(outGoing); err != nil {
		pluto.ApplicationLogger.Error(pluto.ApplicationLog{
			Message: "Stream encoder failed",
			Extra: map[string]any{
				"issuer": ProcessorName_SendResponse,
				"error":  err.Error(),
			},
		})
		return processable, false
	}

	return processable, true
}
