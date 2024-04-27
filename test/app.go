package test

import (
	"database/sql"

	"github.com/axseem/learway/core"
	"github.com/axseem/learway/storage/sqlite"
	_ "modernc.org/sqlite"
)

func App() (*core.App, func()) {
	sqliteDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	sqlite.ForceMigrate(sqliteDB)
	queries := sqlite.New(sqliteDB)
	return core.NewApp(queries), func() { sqliteDB.Close() }
}
