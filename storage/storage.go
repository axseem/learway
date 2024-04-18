// Package storage defines abstractions that data storage must implement.
package storage

type Queries struct {
	Deck    DeckRepo
	Session SessionRepo
	User    UserRepo
}
