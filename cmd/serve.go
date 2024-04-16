package cmd

import (
	"database/sql"

	"github.com/axseem/learway/api"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/axseem/learway/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve() error {
	dbFileName := "learway"
	sqliteDB, err := sql.Open("sqlite", "./"+dbFileName+".db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()
	queries := sqlite.New(sqliteDB)

	if err = sqlite.Migrate(sqliteDB, dbFileName); err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Decompress())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	api.API(e, queries)
	web.App(e)

	e.Logger.Fatal(e.Start(":1323"))
	return nil
}
