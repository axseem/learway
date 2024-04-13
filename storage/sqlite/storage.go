package sqlite

import (
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/storage/sqlite/sqlc"
)

func New(db sqlc.DBTX) *storage.Queries {
	queries := sqlc.New(db)

	return &storage.Queries{
		Deck:    NewDeckStorage(queries),
		Session: NewSessionStorage(queries),
		User:    NewUserStorage(queries),
	}
}
