package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:build
var buildFiles embed.FS

func buildFS() fs.FS {
	return echo.MustSubFS(buildFiles, "build")
}

func App(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(buildFS()),
		HTML5:      true,
	}))
}
