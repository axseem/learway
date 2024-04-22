// Package cmd contains app execution scenarios.
package cmd

import (
	"database/sql"

	"github.com/axseem/learway/api"
	"github.com/axseem/learway/core"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/axseem/learway/web"
)

func Serve() error {
	sqliteDB, err := sql.Open("sqlite", "./learway.db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()
	queries := sqlite.New(sqliteDB)

	app := core.NewApp(queries)
	api.Attach(app)
	web.App(app.Echo)

	return app.Serve(":1323")
}
