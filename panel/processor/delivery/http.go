package delivery

import (
	"net/http"
	"pluto"
	"pluto/panel/pkg/wrapper"
	"pluto/panel/processor"
)

func init() {
	pluto.HTTPServer.GET("/processors",
		wrapper.New[wrapper.EmptyRequest](func(_ wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			return writer.JSON(http.StatusOK, processor.Processors)
		}).Handle(),
	)
}
