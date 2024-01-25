package delivery

import (
	"net/http"
	"pluto/panel/database"
	"pluto/panel/delivery"
	"pluto/panel/extension"
	"pluto/panel/extension/delivery/controller"
	"pluto/panel/pkg/wrapper"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func init() {
	v1 := delivery.HTTPServer.Group("/api/v1", echojwt.WithConfig(delivery.DefaultJWTConfig))

	v1.POST("/extensions/:extension_id",
		wrapper.New[controller.InstallExtension](func(c controller.InstallExtension, w wrapper.ResponseWriter) error {
			return c.Exec(w)
		}).Handle(),
	)

	v1.DELETE("/extensions/:extension_id",
		wrapper.New[controller.UninstallExtension](func(c controller.UninstallExtension, w wrapper.ResponseWriter) error {
			return c.Exec(w)
		}).Handle(),
	)

	v1.GET("/extensions",
		wrapper.New[wrapper.EmptyRequest](func(finder wrapper.EmptyRequest, writer wrapper.ResponseWriter) error {
			tx, err := database.Get().NewTransaction(false)
			if err != nil {
				return wrapper.WriteError(err, writer)
			}

			extensions := []extension.Extension{}
			for _, descriptor := range extension.Descriptors {
				d, _ := extension.FindInstallationDetail(tx, descriptor.ID)
				extensions = append(extensions, extension.Extension{
					Descriptor:         descriptor,
					InstallationDetail: d,
				})
			}

			_ = tx.CommitOrRollback()
			return writer.JSON(http.StatusOK, extensions)
		}).Handle(),
	)
}
