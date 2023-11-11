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
func Process(processable RoutableProcessable) {
	if processable.GetConsumer().PredefinedKind() != KindPipeline {
		ApplicationLogger.Debug(ApplicationLog{
			Message: "Kind is not supported for processing OutComingProcessable",
			Extra:   map[string]any{"kind": processable.GetConsumer().PredefinedKind()},
		})
		return
	}

	executionCacheMutex.RLock()
	defer executionCacheMutex.RUnlock()

	p, found := executionCache[processable.GetConsumer().UniqueProperty()]
	if !found {
		ApplicationLogger.Warning(ApplicationLog{
			Message: "Pipeline not found",
			Extra:   map[string]any{"unique_property": processable.GetConsumer().UniqueProperty()},
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
