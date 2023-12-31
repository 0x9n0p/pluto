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
	e.Use(middleware.CORS())

	if Env.Debug {
		e.Use(middleware.Logger())
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			host, found := HTTPHosts[c.Request().Host]
			if found {
				host.ServeHTTP(c.Response(), c.Request())
			} else {
				err = echo.ErrNotFound
			}
			return
		}
	})

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
	RegisterHTTPHost("", root)
	RegisterHTTPHost("www", root)

	panel := echo.New()
	panel.Use(middleware.Recover())
	panel.Use(middleware.CORSWithConfig(cors))
	RegisterHTTPHost("panel", panel)

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
		if err := HTTPServer.StartTLS(Env.HTTPServerAddress, "/etc/ssl/certs/plutoengine.crt", "/etc/ssl/private/plutoengine.key"); err != nil {
			Log.Fatal("Running HTTP admin server", zap.Error(err))
		}
	}()
}

func RegisterHTTPHost(subdomain string, e *echo.Echo) {
	for _, host := range Env.Host {
		HTTPHosts[MakeHostName(subdomain, host)] = e
	}
}

func FindHTTPHost(subdomain string) *echo.Echo {
	for _, host := range Env.Host {
		e, found := HTTPHosts[MakeHostName(subdomain, host)]
		if found {
			return e
		}
	}
	return nil
}

func MakeHostName(subdomain string, host string) string {
	if subdomain != "" {
		subdomain += "."
	}
	return fmt.Sprintf("%s%s", subdomain, host)
}
