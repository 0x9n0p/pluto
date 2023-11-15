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

	//for {
	//	p, err := client.Receive()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	t.Log(p)
	//}
}
