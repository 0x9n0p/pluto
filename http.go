package pluto

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var HTTPHosts = map[string]*echo.Echo{}

var HTTPServer = func() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())

	if Env.Debug {
		e.Use(middleware.Logger())
	}

	e.Any("/*", func(c echo.Context) (err error) {
		host, found := HTTPHosts[c.Request().Host]
		if found {
			host.ServeHTTP(c.Response(), c.Request())
		} else {
			err = echo.ErrNotFound
		}
		return
	})

	return e
}()

func init() {
	root := echo.New()
	root.Use(middleware.Recover())
	HTTPHosts[Env.HOST] = root
	HTTPHosts[MakeHost("www")] = root

	panel := echo.New()
	panel.Use(middleware.Recover())
	HTTPHosts[MakeHost("panel")] = panel

	if Env.Debug {
		root.Use(middleware.Logger())
		panel.Use(middleware.Logger())
	}

	root.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PlutoEngine.")
	})

	root.POST("/reload", func(ctx echo.Context) error {
		reloadConfig := ctx.QueryParam("config") == "true"

		if reloadConfig {
			var cfg Config
			if err := ctx.Bind(&cfg); err != nil {
				return ctx.String(http.StatusBadRequest, err.Error())
			}

			ReloadExecutionCache(ResolveConfig(cfg))
		}

		return ctx.NoContent(http.StatusOK)
	})

	go func() {
		if err := HTTPServer.Start(Env.HTTPServerAddress); err != nil {
			Log.Fatal("Running HTTP admin server", zap.Error(err))
		}
	}()
}

func MakeHost(subdomain string) string {
	return fmt.Sprintf("%s.%s", subdomain, Env.HOST)
}
