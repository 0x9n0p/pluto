package pluto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// ApplicationLogger
//
// TODO:
//  1. Authorization for ApplicationLogger is required.
//  2. Add other log levels, error, info ..
var ApplicationLogger = func() ApplicationLogCollector {
	logger := ApplicationLogCollector{NewChannel("APPLICATION_LOGGER", 10)}
	ChannelsMutex.Lock()
	defer ChannelsMutex.Unlock()
	Channels[logger.Channel.ID] = logger.Channel
	return logger
}()

type ApplicationLogCollector struct {
	Channel Channel
}

type ApplicationLog struct {
	Message   string         `json:"message"`
	Extra     map[string]any `json:"extra,omitempty"`
	Level     string         `json:"level"`
	CreatedAt time.Time      `json:"created_at"`
}

func (l *ApplicationLogCollector) Debug(log ApplicationLog) {
	log.Level = "Debug"
	l.Log(log)
}

func (l *ApplicationLogCollector) Warning(log ApplicationLog) {
	log.Level = "Warning"
	l.Log(log)
}

func (l *ApplicationLogCollector) Error(log ApplicationLog) {
	log.Level = "Error"
	l.Log(log)
}

func (l *ApplicationLogCollector) Log(log ApplicationLog) {
	log.CreatedAt = time.Now()

	// TODO: Do not convert log to bytes. The subscriber may do it.
	b, err := json.Marshal(log)
	if err != nil {
		Log.Error("Marshalling ApplicationLog", zap.Error(err))
		return
	}

	l.Channel.Publish(&InternalProcessable{
		ID:        uuid.New(), // TODO: Do not generate uuid every time!
		Body:      b,
		CreatedAt: time.Now(),
	})
}
