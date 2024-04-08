package sqlite

import (
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/storage/sqlite/sqlc"
)

func New(db sqlc.DBTX) *storage.Queries {
	return &storage.Queries{
		Deck:    NewDeckStorage(db),
		Session: NewSessionStorage(db),
		User:    NewUserStorage(db),
	}
}
