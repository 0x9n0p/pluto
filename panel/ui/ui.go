package ui

import (
	"embed"
	"pluto/panel/delivery"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:out
	UI          embed.FS
	UIPublicDir = echo.MustSubFS(UI, "out")

	//go:embed out/index.html
	PageHome embed.FS

	//go:embed out/404.html
	PageNotFound embed.FS

	//go:embed out/auth.html
	PageAuth embed.FS

	//go:embed out/logsview.html
	PageLogsView embed.FS

	//go:embed out/pipelines.html
	PagePipelines embed.FS

	//go:embed out/pipelines/create.html
	PagePipelinesCreate embed.FS

	//go:embed out/pipelines/edit.html
	PagePipelinesEdit embed.FS
)

func init() {
	delivery.HTTPServer.StaticFS("/", UIPublicDir)
	delivery.HTTPServer.FileFS("/", "index.html", echo.MustSubFS(PageHome, "out"))
	delivery.HTTPServer.FileFS("/404", "404.html", echo.MustSubFS(PageNotFound, "out"))
	delivery.HTTPServer.FileFS("/auth", "auth.html", echo.MustSubFS(PageAuth, "out"))
	delivery.HTTPServer.FileFS("/logsview", "logsview.html", echo.MustSubFS(PageLogsView, "out"))
	delivery.HTTPServer.FileFS("/pipelines", "pipelines.html", echo.MustSubFS(PagePipelines, "out"))
	delivery.HTTPServer.FileFS("/pipelines/create", "pipelines/create.html", echo.MustSubFS(PagePipelinesCreate, "out"))
	delivery.HTTPServer.FileFS("/pipelines/edit", "pipelines/edit.html", echo.MustSubFS(PagePipelinesEdit, "out"))
}
