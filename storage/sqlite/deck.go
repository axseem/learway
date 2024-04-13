package sqlite

import (
	"context"
	"encoding/json"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/storage/sqlite/sqlc"
)

type DeckStorage struct {
	queries *sqlc.Queries
}

func NewDeckStorage(queries *sqlc.Queries) *DeckStorage {
	return &DeckStorage{queries: queries}
}

func (s DeckStorage) Get(ctx context.Context, id string) (model.Deck, error) {
	sqlcDeck, err := s.queries.GetDeck(ctx, id)
	if err != nil {
		return model.Deck{}, err
	}

	modelDeck := model.Deck{
		ID:        sqlcDeck.ID,
		Title:     sqlcDeck.Title,
		CreatedAt: sqlcDeck.CreatedAt,
		UpdatedAt: sqlcDeck.UpdatedAt,
	}

	if err := json.Unmarshal(sqlcDeck.Cards, &modelDeck.Cards); err != nil {
		return model.Deck{}, err
	}

	return modelDeck, nil
}

func (s DeckStorage) List(ctx context.Context) ([]model.Deck, error) {
	sqlcDecks, err := s.queries.ListDecks(ctx)
	if err != nil {
		return nil, err
	}

	var modelDecks []model.Deck
	for _, sqlcDeck := range sqlcDecks {
		modelDeck := model.Deck{
			ID:        sqlcDeck.ID,
			Title:     sqlcDeck.Title,
			CreatedAt: sqlcDeck.CreatedAt,
			UpdatedAt: sqlcDeck.UpdatedAt,
		}

		if err := json.Unmarshal(sqlcDeck.Cards, &modelDeck.Cards); err != nil {
			return nil, err
		}

		modelDecks = append(modelDecks, modelDeck)
	}

	return modelDecks, nil
}

func (s *DeckStorage) Create(ctx context.Context, arg storage.DeckCreateParams) error {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.queries.CreateDeck(ctx, sqlc.CreateDeckParams{
		ID:    arg.ID,
		Title: arg.Title,
		Cards: cardsJSON,
	})
}

func (s *DeckStorage) Update(ctx context.Context, id string, arg storage.DeckUpdateParams) error {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.queries.UpdateDeck(ctx, sqlc.UpdateDeckParams{
		ID:    id,
		Title: arg.Title,
		Cards: cardsJSON,
	})
}

func (s *DeckStorage) Delete(ctx context.Context, id string) error {
	return s.queries.DeleteDeck(ctx, id)
}
