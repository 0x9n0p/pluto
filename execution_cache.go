package pluto

import (
	"sync"
)

var (
	executionCache      = make(map[string]Pipeline)
	executionCacheMutex = new(sync.RWMutex)
)

// Process
//
// TODO:
//  1. Goroutine pool
func Process(processable Processable) {
	executionCacheMutex.RLock()
	defer executionCacheMutex.RUnlock()

	p, found := executionCache[processable.Pipeline]
	if !found {
		ApplicationLogger.Warning(ApplicationLog{
			Message: "Pipeline not found",
			Extra:   map[string]any{"pipeline": processable.Pipeline},
		})
		return
	}

	p.Process(processable)
}

func ReloadExecutionCache(c map[string]Pipeline) {
	executionCacheMutex.Lock()
	defer executionCacheMutex.Unlock()
	executionCache = c
}
