package delivery

import (
	"net/http"
	"pluto/panel/delivery"
	"pluto/panel/extension"
	"pluto/panel/pkg/wrapper"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := delivery.HTTPServer.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.GET("/extensions",
		wrapper.New[wrapper.EmptyRequest](func(finder wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			return writer.JSON(http.StatusOK, extension.GetStorage().Descriptors)
		}).Handle(),
	)
}
