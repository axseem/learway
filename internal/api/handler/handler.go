package handler

import (
	"github.com/axseem/learway/internal/model"
)

type BaseHandeler struct {
	DeckService model.DeckRepo
}

func NewBaseHandler(ds model.DeckRepo) *BaseHandeler {
	return &BaseHandeler{
		DeckService: ds,
	}
}
