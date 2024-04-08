package cmd

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/axseem/learway/api"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/axseem/learway/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func Serve(port string) error {
	if port == "" {
		port = "1323"
	} else {
		p, err := strconv.Atoi(port)
		if err != nil || p < 1 || p > 65535 {
			return errors.New("invalid port")
		}
	}

	sqliteDB, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()

	queries := sqlite.New(sqliteDB)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	api.API(e, queries)
	web.App(e)
	e.Logger.Fatal(e.Start(":" + port))
	return nil
}
