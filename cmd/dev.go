package cmd

import (
	"database/sql"

	"github.com/axseem/learway/api"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/labstack/echo/v4"
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

	if err = sqlite.Migrate(sqliteDB); err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	api.API(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
	return nil
}
