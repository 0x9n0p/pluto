package delivery

import (
	"net/http"
	"pluto/panel/delivery"
	"pluto/panel/pkg/wrapper"
	"pluto/panel/processor"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := delivery.HTTPServer.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.GET("/processors",
		wrapper.New[processor.DescriptorFinder](func(finder processor.DescriptorFinder, writer wrapper.ResponseWriter) error {
			return writer.JSON(http.StatusOK, finder.Find())
		}).Handle(),
	)
}
