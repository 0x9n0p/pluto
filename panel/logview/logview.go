package logview

import (
	"io"
	"pluto"
)

type BindApplicationLogs struct {
	Identifier pluto.Identifier
	Writer     io.Writer
}

func (b *BindApplicationLogs) Bind() {
	pluto.ApplicationLogger.Channel.Join(pluto.BaseJoinable{
		Identifier: b.Identifier,
		Processor: pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
			p, ok := (&pluto.WriteToIOProcessor{Writer: b.Writer}).Process(processable)
			if !ok {
				pluto.ApplicationLogger.Channel.Leave(b.Identifier)
				return p, false
			}
			return p, true
		}),
	})
}
