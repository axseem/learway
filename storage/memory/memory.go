// Package memory implements in memory storage mainly for easier testing of the app logic.
package memory

import (
	"github.com/axseem/learway/storage"
)

// New creates a new in memory storage queries instance.
func New() *storage.Queries {
	return &storage.Queries{
		Deck:    &DeckStorage{},
		Session: &SessionStorage{},
		User:    &UserStorage{},
	}
}
