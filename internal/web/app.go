package web

import (
	"embed"
	"io/fs"

	"github.com/axseem/learway/internal/database"
	"github.com/labstack/echo/v4"
)

//go:embed all:build
var buildFiles embed.FS

func buildFS() fs.FS {
	return echo.MustSubFS(buildFiles, "build")
}

func App(e *echo.Echo, db *database.Queries) {
	e.StaticFS("/", buildFS())
}
