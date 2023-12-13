package pluto

import (
	"encoding/json"
)

const ProcessorName_SendResponse = "SEND_RESPONSE"

func init() {
	PredefinedProcessors[ProcessorName_SendResponse] = func(args []Value) (p Processor, err error) {
		defer creatorPanicHandler(ProcessorName_SendResponse, &err)()
		return SendResponse{
			PipelineName: Find("pipeline_name", args...).Get().(string),
		}, err
	}
}

// SendResponse Encodes the body, sends it back to the producer and never changes the input processable.
type SendResponse struct {
	// TODO
	// Filter []string
	PipelineName string
}

func (p SendResponse) Process(processable Processable) (Processable, bool) {
	outGoing := OutGoingProcessable{
		Consumer: ExternalIdentifier{
			Name: p.PipelineName,
			Kind: KindPipeline,
		},
		Body: processable.GetBody(),
	}

	outComing, ok := processable.(*OutComingProcessable)
	if !ok {
		ApplicationLogger.Error(ApplicationLog{
			Message: "The processable is not an out_coming_processable",
			Extra:   map[string]any{"issuer": ProcessorName_SendResponse},
		})
		return processable, false
	}

	b, err := json.Marshal(outGoing)
	if err != nil {
		ApplicationLogger.Error(ApplicationLog{
			Message: "Failed to encode processable",
			Extra: map[string]any{
				"issuer": ProcessorName_SendResponse,
				"error":  err.Error(),
			},
		})
		return processable, false
	}

	_, err = outComing.Connection.Conn.Write(b)
	if err != nil {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Failed to write to connection",
			Extra: map[string]any{
				"issuer": ProcessorName_SendResponse,
				"error":  err.Error(),
			},
		})
		return processable, false
	}

	return processable, true
}
