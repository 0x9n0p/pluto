package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"pluto"
	"pluto/panel/database"
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

	v1 := delivery.HTTPServer.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.GET("/pipelines",
		wrapper.New[wrapper.EmptyRequest](func(request wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			tx, err := database.Get().NewTransaction(false)
			if err != nil {
				return wrapper.WriteError(err, writer)
			}

			list, err := pipeline.All(tx)
			if err != nil {
				_ = tx.Rollback()
				return WriteError(err, writer)
			}

			if err := tx.Rollback(); err != nil {
				return wrapper.WriteError(err, writer)
			}

			return writer.JSON(http.StatusOK, list)
		}).Handle(),
	)

	v1.POST("/pipelines",
		wrapper.New[pipeline.Pipeline](func(p pipeline.Pipeline, writer wrapper.ResponseWriter) error {
			tx, err := database.Get().NewTransaction(true)
			if err != nil {
				return wrapper.WriteError(err, writer)
			}

			p.Transaction = tx

			if err := p.Save(); err != nil {
				_ = tx.Rollback()
				return WriteError(err, writer)
			}

			if err := tx.CommitOrRollback(); err != nil {
				return WriteError(err, writer)
			}

			if err := pipeline.ReloadExecutionCache(); err != nil {
				return WriteError(fmt.Errorf("reload execution cache: %v", err), writer)
			}

			return writer.JSON(http.StatusCreated, p)
		}).Handle(),
	)

	// TODO: Add a controller
	v1.DELETE("/pipelines",
		wrapper.New[struct {
			Name string `query:"name" validate:"required"`
		}](func(r struct {
			Name string `query:"name" validate:"required"`
		}, writer wrapper.ResponseWriter) error {
			tx, err := database.Get().NewTransaction(true)
			if err != nil {
				return wrapper.WriteError(err, writer)
			}

			p, err := pipeline.Find(tx, r.Name)
			if err != nil {
				_ = tx.Rollback()
				return WriteError(err, writer)
			}

			if err := p.Delete(); err != nil {
				_ = tx.Rollback()
				return WriteError(err, writer)
			}

			if err := tx.CommitOrRollback(); err != nil {
				return wrapper.WriteError(err, writer)
			}

			// TODO
			if err := pipeline.ReloadExecutionCache(); err != nil {
				_ = tx.Rollback()
				return WriteError(fmt.Errorf("reload execution cache: %v", err), writer)
			}

			return writer.JSON(http.StatusOK, p)
		}).Handle(),
	)
}
