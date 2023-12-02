package delivery

import (
	"net/http"
	"pluto"
	"pluto/panel/pkg/wrapper"
	"pluto/panel/processor"
)

func init() {
	pluto.FindHTTPHost("panel").GET("/api/v1/processors",
		wrapper.New[processor.DescriptorFinder](func(finder processor.DescriptorFinder, writer wrapper.ResponseWriter) error {
			return writer.JSON(http.StatusOK, finder.Find())
		}).Handle(),
	)
}
