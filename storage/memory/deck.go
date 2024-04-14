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

func (s DeckStorage) GetByUserID(ctx context.Context, userID string) ([]model.Deck, error) {
	var decks []model.Deck

	for _, deck := range s {
		if deck.UserID == userID {
			decks = append(decks, deck)
		}
	}

	if len(decks) == 0 {
		return nil, ErrNotFound
	}

	return decks, nil
}

func (s DeckStorage) List(ctx context.Context) ([]model.Deck, error) {
	return s, nil
}

func (s *DeckStorage) Create(ctx context.Context, arg storage.DeckCreateParams) error {
	newDeck := model.Deck{
		ID:        arg.ID,
		UserID:    arg.UserID,
		Title:     arg.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Cards:     arg.Cards,
	}
	*s = append(*s, newDeck)
	return nil
}

func (s *DeckStorage) Update(ctx context.Context, id string, arg storage.DeckUpdateParams) error {
	for i, deck := range *s {
		if deck.ID == id {
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
