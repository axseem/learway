package storage

import (
	"context"

	"github.com/axseem/learway/model"
)

type DeckCreateParams struct {
	ID     string
	UserID string
	Title  string
	Cards  model.Cards
}

type DeckUpdateParams struct {
	Title string
	Cards model.Cards
}

type DeckRepo interface {
	Get(ctx context.Context, id string) (model.Deck, error)
	List(ctx context.Context) ([]model.Deck, error)
	Create(ctx context.Context, arg DeckCreateParams) error
	Update(ctx context.Context, id string, arg DeckUpdateParams) error
	Delete(ctx context.Context, id string) error
}
