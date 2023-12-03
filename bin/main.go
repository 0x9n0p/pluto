package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pluto"
	_ "pluto/panel"
	"pluto/panel/pipeline"
	"pluto/panel/processor"
)

func init() {
	p := pipeline.Pipeline{
		Name: "my pipeline",
		Processors: []processor.Processor{
			{
				Name: pluto.ProcessorName_ExecAndJoinResult,
				Arguments: []pluto.Value{
					{
						Name: "Processor",
						Type: pluto.TypeProcessor,
						Value: map[string]any{
							"name": pluto.ProcessorName_WriteToInputOutput,
							"arguments": []pluto.Value{
								{
									Name:  "inc",
									Type:  "df",
									Value: "asdf",
								},
							},
						},
					},
				},
			},
		},
	}

	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	pluto.ApplicationLogger.Channel.Join(pluto.BaseJoinable{
		Identifier: pluto.ExternalIdentifier{
			Name: "STD_OUT",
			Kind: "LocalStream",
		},
		Processor: pluto.WriteToIOProcessor{Writer: os.Stdout},
	})
}

func main() {
	fmt.Printf("%s %s\n", pluto.Name, pluto.Version)
	select {}
}
