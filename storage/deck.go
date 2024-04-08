package storage

import (
	"context"

	"github.com/axseem/learway/model"
)

type DeckCreateParams struct {
	ID    string
	Title string
	Cards [][2]string
}

type DeckQueries interface {
	Get(ctx context.Context, id string) (model.Deck, error)
	List(ctx context.Context) ([]model.Deck, error)
	Create(ctx context.Context, arg DeckCreateParams) error
	Update(ctx context.Context, arg DeckCreateParams) error
	Delete(ctx context.Context, id string) error
}
