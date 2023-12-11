package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/delivery"
	"pluto/panel/pipeline"
	"pluto/panel/pkg/wrapper"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	WriteError := func(err error, writer wrapper.ResponseWriter) error {
		var perr *pluto.Error
		if errors.As(err, &perr) {
			return writer.JSON(perr.HTTPCode, perr)
		}
		return writer.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	panel := pluto.FindHTTPHost(pluto.PanelSubdomain)
	v1 := panel.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.GET("/pipelines",
		wrapper.New[wrapper.EmptyRequest](func(request wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			list, err := pipeline.GetStorage().List()
			if err != nil {
				return WriteError(err, writer)
			}
			return writer.JSON(http.StatusOK, list)
		}).Handle(),
	)

	v1.POST("/pipelines",
		wrapper.New[pipeline.Pipeline](func(p pipeline.Pipeline, writer wrapper.ResponseWriter) error {
			if err := p.Save(); err != nil {
				return WriteError(err, writer)
			}

			if err := pipeline.GetStorage().ReloadExecutionCache(); err != nil {
				return WriteError(fmt.Errorf("reload execution cache: %v", err), writer)
			}

			return writer.JSON(http.StatusCreated, p)
		}).Handle(),
	)

	v1.DELETE("/pipelines",
		wrapper.New[struct {
			Name string `query:"name" validate:"required"`
		}](func(r struct {
			Name string `query:"name" validate:"required"`
		}, writer wrapper.ResponseWriter) error {
			p, err := pipeline.GetStorage().Find(r.Name)
			if err != nil {
				return WriteError(err, writer)
			}

			if err := p.Delete(); err != nil {
				return WriteError(err, writer)
			}

			return writer.JSON(http.StatusOK, p)
		}).Handle(),
	)
}
