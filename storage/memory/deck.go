package memory

import (
	"context"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
)

type DeckStorage []model.Deck

func (s DeckStorage) Get(ctx context.Context, id string) (model.Deck, error) {
	for _, deck := range s {
		if deck.ID == id {
			return deck, nil
		}
	}
	return model.Deck{}, ErrNotFound
}

func (s DeckStorage) List(ctx context.Context) ([]model.Deck, error) {
	return s, nil
}

func (s *DeckStorage) Create(ctx context.Context, arg storage.DeckCreateParams) error {
	newDeck := model.Deck{
		ID:        arg.ID,
		Title:     arg.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Cards:     arg.Cards,
	}
	*s = append(*s, newDeck)
	return nil
}

func (s *DeckStorage) Update(ctx context.Context, arg storage.DeckCreateParams) error {
	for i, deck := range *s {
		if deck.ID == arg.ID {
			(*s)[i].Title = arg.Title
			(*s)[i].Cards = arg.Cards
			return nil
		}
	}
	return ErrNotFound
}

func (s *DeckStorage) Delete(ctx context.Context, id string) error {
	for i, deck := range *s {
		if deck.ID == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
