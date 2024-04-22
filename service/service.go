package service

import (
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
)

type Base struct {
	Deck    *Deck
	Session *Session
	User    *User
}

func NewBase(queries *storage.Queries, validator *validator.Validate) *Base {
	deckService := NewDeckService(queries, validator)
	userService := NewUserService(queries, validator)
	sessionService := NewSessionService(queries, validator)

	return &Base{
		Deck:    deckService,
		Session: sessionService,
		User:    userService,
	}
}
