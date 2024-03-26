package service

import (
	"context"
	"encoding/json"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type DeckService struct {
	db *database.Queries
}

func NewDeckService(db *database.Queries) *DeckService {
	return &DeckService{
		db: db,
	}
}

func (s DeckService) Get(ctx context.Context, id string) (*model.Deck, error) {
	raw, err := s.db.GetDeck(ctx, id)
	if err != nil {
		return &model.Deck{}, err
	}

	deck := &model.Deck{
		ID:        raw.ID,
		Title:     raw.Title,
		CreatedAt: raw.CreatedAt,
		UpdatedAt: raw.UpdatedAt,
	}

	if err := json.Unmarshal(raw.Cards, &deck.Cards); err != nil {
		return &model.Deck{}, err
	}

	return deck, nil
}

func (s DeckService) List(ctx context.Context) ([]model.Deck, error) {
	rawList, err := s.db.ListDecks(ctx)
	if err != nil {
		return []model.Deck{}, err
	}

	var deckList []model.Deck

	for _, raw := range rawList {
		deck := model.Deck{
			ID:        raw.ID,
			Title:     raw.Title,
			CreatedAt: raw.CreatedAt,
			UpdatedAt: raw.UpdatedAt,
		}

		if err := json.Unmarshal(raw.Cards, &deck.Cards); err != nil {
			return []model.Deck{}, err
		}

		deckList = append(deckList, deck)
	}

	return deckList, nil
}

func (s DeckService) Create(ctx context.Context, arg model.DeckCreateParams) error {
	id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		return err
	}

	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.db.CreateDeck(ctx, database.CreateDeckParams{
		ID:    id,
		Title: arg.Title,
		Cards: cardsJSON,
	})
}

func (s DeckService) Update(ctx context.Context, arg model.DeckUpdateParams) error {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.db.UpdateDeck(ctx, database.UpdateDeckParams{
		ID:    arg.ID,
		Title: arg.Title,
		Cards: cardsJSON,
	})
}

func (s DeckService) Delete(ctx context.Context, id string) error {
	return s.db.DeleteDeck(ctx, id)
}
