package panel

import (
	"pluto"
	"pluto/panel/pipeline"

	_ "pluto/panel/database"
	_ "pluto/panel/delivery"

	_ "pluto/panel/account/delivery"
	_ "pluto/panel/extension/delivery"
	_ "pluto/panel/logview/delivery"
	_ "pluto/panel/pipeline/delivery"
	_ "pluto/panel/processor/delivery"
	_ "pluto/panel/statistics/delivery"

	_ "pluto/panel/ui"

	"go.uber.org/zap"
)

func init() {
	if err := pipeline.ReloadExecutionCache(); err != nil {
		pluto.Log.Fatal("Reload execution cache", zap.Error(err))
	}
}
