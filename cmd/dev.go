package cmd

import (
	"database/sql"

	"github.com/axseem/learway/api"
	"github.com/axseem/learway/core"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func Dev() error {
	sqliteDB, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()
	queries := sqlite.New(sqliteDB)

	app := core.NewApp(queries)
	api.Attach(app)

	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	return app.Serve(":1323")
}
