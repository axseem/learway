package memory

import (
	"github.com/axseem/learway/storage"
)

func New() *storage.Queries {
	return &storage.Queries{
		Deck:    &DeckStorage{},
		Session: &SessionStorage{},
		User:    &UserStorage{},
	}
}
