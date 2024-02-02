package pluto

import (
	"time"

	"github.com/google/uuid"
)

// OnInitialized is called when all the subsystems are initialized and any pipeline can be executed.
// The pipelines that has field Pipeline.ExecuteOnStartup enabled will be executed.
func OnInitialized() {
	ExecutionCacheMutex.RLock()
	defer ExecutionCacheMutex.RUnlock()

	for _, pipeline := range ExecutionCache {
		if !pipeline.ExecuteOnStartup {
			continue
		}

		pipeline.Process(&InternalProcessable{
			ID:        uuid.New(),
			Body:      map[string]any{},
			CreatedAt: time.Now(),
		})
	}
}
