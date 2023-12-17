package sdk_test

import (
	"pluto/sdk"
	"testing"
)

func TestTCP(t *testing.T) {
	client := sdk.TCPClient{
		ServerAddress:      "localhost:9631",
		Producer:           sdk.Identifier{},
		ProducerCredential: sdk.Credential{},
	}

	if err := client.Connect(); err != nil {
		t.Fatal(err)
	}

	go func() {
		for i := 0; i < 200000; i++ {
			go func(i int) {
				err := client.Send(sdk.Processable{
					Consumer: sdk.Identifier{
						Name: "PING",
						Kind: "Pipeline",
					},
					Body: i,
				})
				if err != nil {
					t.Fatal(err)
					return
				}
			}(i)
		}
	}()

	for {
		p, err := client.Receive()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(p)
	}
}
