package model

import (
	"context"
	"time"
)

type Deck struct {
	ID        string
	Title     string
	Cards     [][2]string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeckCreateParams struct {
	Title string
	Cards [][2]string
}

type DeckUpdateParams struct {
	ID    string
	Title string
	Cards [][2]string
}

type DeckRepo interface {
	Get(ctx context.Context, id string) (*Deck, error)
	List(ctx context.Context) ([]Deck, error)
	Create(ctx context.Context, arg DeckCreateParams) (*Deck, error)
	Update(ctx context.Context, arg DeckUpdateParams) error
	Delete(ctx context.Context, id string) error
	ValidateCreateParams(arg DeckCreateParams) error
	ValidateUpdateParams(ctx context.Context, arg DeckUpdateParams) error
}
