package model

import (
	"context"
	"time"
)

type Deck struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Cards     [][2]string `json:"cards"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type DeckCreateParams struct {
	Title string      `json:"title" validate:"required,gt=0,lte=256"`
	Cards [][2]string `json:"cards" validate:"required,gt=0,lte=4096,dive,required,dive,required,gt=0,lte=1024"`
}

type DeckRepo interface {
	Get(ctx context.Context, id string) (Deck, error)
	List(ctx context.Context) ([]Deck, error)
	Create(ctx context.Context, arg DeckCreateParams) (Deck, error)
	Update(ctx context.Context, id string, arg DeckCreateParams) (Deck, error)
	Delete(ctx context.Context, id string) error
}
