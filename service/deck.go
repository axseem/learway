package service

import (
	"context"
	"log"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type DeckService struct {
	storage *storage.Queries
}

func NewDeckService(store *storage.Queries) *DeckService {
	return &DeckService{
		storage: store,
	}
}

func (s DeckService) Get(ctx context.Context, id string) (model.Deck, error) {
	return s.storage.Deck.Get(ctx, id)
}

func (s DeckService) List(ctx context.Context) ([]model.Deck, error) {
	return s.storage.Deck.List(ctx)
}

func (s DeckService) Create(ctx context.Context, arg model.DeckCreateParams) (model.Deck, error) {
	id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		log.Fatal(err)
	}

	deckCreateParams := storage.DeckCreateParams{
		ID:    id,
		Title: arg.Title,
		Cards: arg.Cards,
	}

	if err = s.storage.Deck.Create(ctx, deckCreateParams); err != nil {
		return model.Deck{}, err
	}

	return s.storage.Deck.Get(ctx, id)
}

func (s DeckService) Update(ctx context.Context, id string, arg model.DeckCreateParams) (model.Deck, error) {
	err := s.storage.Deck.Update(ctx, storage.DeckCreateParams{
		ID:    id,
		Title: arg.Title,
		Cards: arg.Cards,
	})
	if err != nil {
		return model.Deck{}, err
	}

	return s.Get(ctx, id)
}

func (s DeckService) Delete(ctx context.Context, id string) error {
	return s.storage.Deck.Delete(ctx, id)
}
