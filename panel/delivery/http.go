package delivery

import (
	"net/http"
	"pluto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var HTTPServer = func() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	if pluto.Env.Debug {
		e.Use(middleware.Logger())
	}

	return e
}()

func init() {
	cors := middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}

	root := echo.New()
	root.Use(middleware.Recover())
	root.Use(middleware.CORSWithConfig(cors))

	if pluto.Env.Debug {
		root.Use(middleware.Logger())
	}

	go func() {
		if err := HTTPServer.StartTLS(pluto.Env.HTTPServerAddress, pluto.Env.HTTPCertificatePath, pluto.Env.HTTPKeyPath); err != nil {
			pluto.Log.Fatal("Panel web server", zap.Error(err))
		}
	}()
}
