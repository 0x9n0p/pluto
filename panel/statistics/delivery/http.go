package delivery

import (
	"net/http"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"
	"pluto/panel/statistics"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := delivery.HTTPServer.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.GET("/statistics",
		wrapper.New[wrapper.EmptyRequest](func(_ wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			return writer.JSON(http.StatusOK, statistics.Get())
		}).Handle(),
	)
}
