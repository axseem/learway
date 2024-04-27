// Package model defines essential types and interfaces for app logic.
package model

import (
	"context"
	"time"
)

type Cards [][2]string

type Deck struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Title     string    `json:"title"`
	Cards     Cards     `json:"cards"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeckCreateParams struct {
	SessionID string `json:"sessionID"`
	Title     string `json:"title" validate:"required,gt=0,lte=256"`
	Cards     Cards  `json:"cards" validate:"required,gt=1,lte=4096,dive,required,dive,required,gt=0,lte=1024"`
	Subject   string `json:"subject" validate:"lte=32"`
}

type DeckUpdateParams struct {
	Title   string `json:"title" validate:"required,gt=0,lte=256"`
	Cards   Cards  `json:"cards" validate:"required,gt=0,lte=4096,dive,required,dive,required,gt=0,lte=1024"`
	Subject string `json:"subject" validate:"lte=32"`
}

type DeckRepo interface {
	Get(ctx context.Context, id string) (Deck, error)
	GetByUsername(ctx context.Context, userID string) ([]Deck, error)
	List(ctx context.Context) ([]Deck, error)
	Create(ctx context.Context, arg DeckCreateParams) (Deck, error)
	Update(ctx context.Context, id string, arg DeckUpdateParams) (Deck, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, title string) ([]Deck, error)
}
