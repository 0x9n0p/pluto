package logview

import (
	"io"
	"pluto"

	"go.uber.org/zap"
)

type BindApplicationLogs struct {
	Identifier pluto.Identifier
	Writer     io.Writer
}

func (b *BindApplicationLogs) Bind() {
	pluto.ApplicationLogger.Channel.Join(pluto.BaseJoinable{
		Identifier: b.Identifier,
		Processor: pluto.NewInlineProcessor(func(processable pluto.Processable) (pluto.Processable, bool) {
			// TODO
			//  https://github.com/gorilla/websocket/issues/119
			//  There most be a read/write queues.
			defer func() {
				if v := recover(); v != nil {
					pluto.Log.Panic("Binding application logs", zap.String("error", v.(string)))
				}
			}()

			p, ok := (&pluto.IOWriter{Writer: b.Writer}).Process(processable)
			if !ok {
				pluto.ApplicationLogger.Channel.Leave(b.Identifier)
				return p, false
			}
			return p, true
		}),
	})
}
