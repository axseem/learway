package cmd

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/axseem/learway/internal/api"
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web"
	"github.com/labstack/echo/v4"
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

	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqlite.Close()

	db := database.New(sqlite)

	e := echo.New()

	api.API(e, db)
	web.App(e, db)
	e.Logger.Fatal(e.Start(":" + port))
	return nil
}
