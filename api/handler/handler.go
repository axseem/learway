package handler

import "github.com/axseem/learway/service"

type Base struct {
	Deck    *Deck
	Session *Session
	User    *User
}

func NewBase(s *service.Base) *Base {
	return &Base{
		Deck:    NewDeck(s.Deck),
		Session: NewSession(s.Session),
		User:    NewUser(s.User),
	}
}
