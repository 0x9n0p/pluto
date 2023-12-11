package panel

import (
	"pluto"
	_ "pluto/panel/auth/delivery"
	_ "pluto/panel/logview/delivery"
	"pluto/panel/pipeline"
	_ "pluto/panel/pipeline/delivery"
	_ "pluto/panel/processor/delivery"
	_ "pluto/panel/statistics/delivery"
	_ "pluto/panel/ui"

	"go.uber.org/zap"
)

func init() {
	if err := pipeline.GetStorage().ReloadExecutionCache(); err != nil {
		pluto.Log.Fatal("Reload execution cache", zap.Error(err))
	}
}
