package handler

import (
	"github.com/axseem/learway/model"
)

type BaseHandeler struct {
	DeckService    model.DeckRepo
	UserService    model.UserRepo
	SessionService model.SessionRepo
}

func NewBaseHandler(d model.DeckRepo, u model.UserRepo, s model.SessionRepo) *BaseHandeler {
	return &BaseHandeler{
		DeckService:    d,
		UserService:    u,
		SessionService: s,
	}
}
