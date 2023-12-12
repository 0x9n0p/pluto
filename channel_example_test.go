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

	ch.Join(&ChannelJoinableProcessor{
		Name: "OUT_1",
		Kind: "STD_IO",
		Processor: pluto.IOWriter{
			Writer: os.Stdout,
		},
	})

	ch.Join(&ChannelJoinableProcessor{
		Name: "OUT_2",
		Kind: "STD_IO",
		Processor: pluto.IOWriter{
			Writer: os.Stdout,
		},
	})

	ch.Join(&ChannelJoinableProcessor{
		Name: "OUT_3",
		Kind: "STD_IO",
		Processor: pluto.IOWriter{
			Writer: os.Stdout,
		},
	})

	ch.Publish(&pluto.InternalProcessable{
		ID:        uuid.New(),
		Body:      []byte("Hello World 1\n"),
		CreatedAt: time.Now(),
	})

	ch.Publish(&pluto.InternalProcessable{
		ID:        uuid.New(),
		Body:      []byte("Hello World 2\n"),
		CreatedAt: time.Now(),
	})

	<-time.Tick(time.Second)
}

// TODO: Remove it and use pluto.BaseJoinable instead
type ChannelJoinableProcessor struct {
	Name string
	Kind string
	pluto.Processor
}

func (c *ChannelJoinableProcessor) UniqueProperty() string {
	return c.Name
}

func (c *ChannelJoinableProcessor) PredefinedKind() string {
	return c.Kind
}
