package storage

type Queries struct {
	Deck    DeckRepo
	Session SessionRepo
	User    UserRepo
}
