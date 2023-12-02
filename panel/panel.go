package panel

import (
	"net/http"
	"pluto"
	"pluto/panel/pipeline"
	_ "pluto/panel/pipeline/delivery"
	_ "pluto/panel/processor/delivery"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func init() {
	pluto.FindHTTPHost("panel").GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Panel.")
	})

	if err := pipeline.GetStorage().ReloadExecutionCache(); err != nil {
		pluto.Log.Fatal("Reload execution cache", zap.Error(err))
	}
}
