package pluto_test

import (
	"os"
	"pluto"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestChannel(t *testing.T) {
	ch := pluto.NewChannel("MY_CHANNEL", 10)

	ch.Join(&pluto.ChannelJoinableProcessor{
		Name: "OUT_1",
		Kind: "STD_IO",
		Processor: pluto.WriteToIOProcessor{
			Writer: os.Stdout,
		},
	})

	ch.Join(&pluto.ChannelJoinableProcessor{
		Name: "OUT_2",
		Kind: "STD_IO",
		Processor: pluto.WriteToIOProcessor{
			Writer: os.Stdout,
		},
	})

	ch.Join(&pluto.ChannelJoinableProcessor{
		Name: "OUT_3",
		Kind: "STD_IO",
		Processor: pluto.WriteToIOProcessor{
			Writer: os.Stdout,
		},
	})

	ch.Publish(&pluto.InternalProcessable{
		ID:        uuid.New(),
		Producer:  TestIdentifier{Name: "TestChannel", Kind: "Test"},
		Body:      []byte("Hello World 1\n"),
		CreatedAt: time.Now(),
	})

	ch.Publish(&pluto.InternalProcessable{
		ID:        uuid.New(),
		Producer:  TestIdentifier{Name: "TestChannel", Kind: "Test"},
		Body:      []byte("Hello World 2\n"),
		CreatedAt: time.Now(),
	})

	<-time.Tick(time.Second)
}
