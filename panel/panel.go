package panel

import (
	"pluto"
	"pluto/panel/pipeline"
	_ "pluto/panel/pipeline/delivery"
	_ "pluto/panel/processor/delivery"

	"go.uber.org/zap"
)

func init() {
	if err := pipeline.GetStorage().ReloadExecutionCache(); err != nil {
		pluto.Log.Fatal("Reload execution cache", zap.Error(err))
	}
}
