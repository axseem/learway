package cmd

import (
	"database/sql"

	"github.com/axseem/learway/storage/sqlite"
)

func Migrate(dbFileName string) error {
	sqliteDB, err := sql.Open("sqlite", "./"+dbFileName+".db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()

	return sqlite.Migrate(sqliteDB, dbFileName)
}
