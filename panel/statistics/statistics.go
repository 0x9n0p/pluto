package statistics

import (
	"pluto"
	"runtime"
	"time"

	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
	"go.uber.org/zap"
)

type Statistics struct {
	ConnectedClients int `json:"connected_clients"`
	WaitingClients   int `json:"waiting_clients"`

	ActivePipelines int `json:"active_pipelines"`

	RunningGoroutines int `json:"running_goroutines"`

	Uptime uint64 `json:"uptime"`

	TotalMemory  uint64 `json:"total_memory"`
	UsedMemory   uint64 `json:"used_memory"`
	CachedMemory uint64 `json:"cached_memory"`
	FreeMemory   uint64 `json:"free_memory"`

	LoadAverage1  float64 `json:"load_average_1"`
	LoadAverage5  float64 `json:"load_average_5"`
	LoadAverage15 float64 `json:"load_average_15"`
}

func Get() (s Statistics) {
	// Connections
	{
		s.ConnectedClients = len(pluto.AuthenticatedConnections)
		s.WaitingClients = len(pluto.AcceptedConnections)
	}

	// Pipelines
	{
		s.ActivePipelines = len(pluto.ExecutionCache)
	}

	// Goroutines
	{
		s.RunningGoroutines = runtime.NumGoroutine()
	}

	// Uptime
	{
		up, err := uptime.Get()
		if err != nil {
			pluto.Log.Error("Cannot get uptime", zap.Error(err))
			up = time.Hour * 999
		}

		s.Uptime = uint64(up.Seconds())
	}

	// Memory
	{
		mem, err := memory.Get()
		if err != nil {
			pluto.Log.Error("Cannot get memory usage", zap.Error(err))
			return
		}

		s.TotalMemory = mem.Total
		s.UsedMemory = mem.Used
		s.CachedMemory = mem.Cached
		s.FreeMemory = mem.Free
	}

	// Load average
	{
		avg, err := loadavg.Get()
		if err != nil {
			pluto.Log.Error("Cannot get load average", zap.Error(err))
			return
		}

		s.LoadAverage1 = avg.Loadavg1
		s.LoadAverage5 = avg.Loadavg5
		s.LoadAverage15 = avg.Loadavg15
	}

	return
}
