package handler

import "github.com/axseem/learway/internal/service"

type BaseHandeler struct {
	DeckService service.DeckService
}

func NewBaseHandler(ds service.DeckService) *BaseHandeler {
	return &BaseHandeler{
		DeckService: ds,
	}
}
