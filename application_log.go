package pluto

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
)

// ApplicationLogger
//
// TODO:
//  1. The output should be changeable to write to network fd
//  2. Add other log levels, error, info ..
var ApplicationLogger = ApplicationLogCollector{Output: os.Stdout}

type ApplicationLogCollector struct {
	Output io.Writer
}

type ApplicationLog struct {
	Message   string         `json:"message"`
	Extra     map[string]any `json:"extra,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
}

func (l *ApplicationLogCollector) Warning(log ApplicationLog) {
	log.CreatedAt = time.Now()

	b, err := json.Marshal(log)
	if err != nil {
		Log.Error("Marshalling ApplicationLog", zap.Error(err))
		return
	}

	_, err = l.Output.Write(b)
	if err != nil {
		Log.Error("Writing ApplicationLog", zap.Error(err))
		return
	}
}
