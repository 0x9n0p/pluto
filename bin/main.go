package main

import (
	"fmt"
	"os"
	"pluto"
	_ "pluto/panel"
	"time"
)

func init() {
	pluto.ApplicationLogger.Channel.Join(pluto.BaseJoinable{
		Identifier: pluto.ExternalIdentifier{
			Name: "STD_OUT",
			Kind: "LocalStream",
		},
		Processor: pluto.WriteToIOProcessor{Writer: os.Stdout},
	})
}

func main() {
	go func() {
		for {
			pluto.ApplicationLogger.Warning(pluto.ApplicationLog{
				Message: "warning log for test",
				Extra:   map[string]any{},
			})
			<-time.Tick(time.Second * 3)
		}
	}()
	fmt.Printf("%s %s\n", pluto.Name, pluto.Version)
	select {}
}
