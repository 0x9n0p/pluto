package pluto

import (
	"sync"
)

var (
	ExecutionCache      = make(map[string]Pipeline)
	ExecutionCacheMutex = new(sync.RWMutex)
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

	ExecutionCacheMutex.RLock()
	defer ExecutionCacheMutex.RUnlock()

	p, found := ExecutionCache[processable.GetConsumer().UniqueProperty()]
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
	ExecutionCacheMutex.Lock()
	defer ExecutionCacheMutex.Unlock()
	ExecutionCache = c
}
