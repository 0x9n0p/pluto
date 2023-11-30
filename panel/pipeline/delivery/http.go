package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/pipeline"
	"pluto/panel/pkg/wrapper"
)

func init() {
	WriteError := func(err error, writer wrapper.ResponseWriter) error {
		var perr *pluto.Error
		if errors.As(err, &perr) {
			return writer.JSON(perr.HTTPCode, perr)
		}
		return writer.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	panel := pluto.FindHTTPHost("panel")

	panel.GET("/pipelines",
		wrapper.New[wrapper.EmptyRequest](func(request wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			list, err := pipeline.GetStorage().List()
			if err != nil {
				return WriteError(err, writer)
			}
			return writer.JSON(http.StatusOK, list)
		}).Handle(),
	)

	panel.POST("/pipelines",
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

	panel.DELETE("/pipelines",
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
