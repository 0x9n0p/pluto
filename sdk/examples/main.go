package main

import (
	"log"
	"pluto/sdk"
	"time"
)

func main() {
	client := sdk.TCPClient{
		ServerAddress:      "localhost:9630",
		Producer:           sdk.Identifier{},
		ProducerCredential: sdk.Credential{},
	}

	if err := client.Connect(); err != nil {
		log.Fatalln(err)
	}

	go func() {
		for i := 0; i < 1; i++ {
			err := client.Send(sdk.Processable{
				Consumer: sdk.Identifier{
					Name: "PING",
					Kind: "Pipeline",
				},
				Body: map[string]any{
					"time": time.Now().UTC().Unix(),
				},
			})
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}()

	for {
		p, err := client.Receive()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(p)
	}
}
