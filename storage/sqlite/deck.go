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

	return convertDeck(sqlcDeck)
}

func (s DeckStorage) GetByUserID(ctx context.Context, userID string) ([]model.Deck, error) {
	sqlcDecks, err := s.queries.GetDeckByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return convertDeckSlice(sqlcDecks)
}

func (s DeckStorage) List(ctx context.Context) ([]model.Deck, error) {
	sqlcDecks, err := s.queries.ListDecks(ctx)
	if err != nil {
		return nil, err
	}
	return convertDeckSlice(sqlcDecks)
}

func (s *DeckStorage) Create(ctx context.Context, arg storage.DeckCreateParams) error {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.queries.CreateDeck(ctx, sqlc.CreateDeckParams{
		ID:      arg.ID,
		UserID:  arg.UserID,
		Title:   arg.Title,
		Cards:   cardsJSON,
		Subject: toNullString(arg.Subject),
	})
}

func (s *DeckStorage) Update(ctx context.Context, id string, arg storage.DeckUpdateParams) error {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return err
	}

	return s.queries.UpdateDeck(ctx, sqlc.UpdateDeckParams{
		ID:      id,
		Title:   arg.Title,
		Cards:   cardsJSON,
		Subject: toNullString(arg.Subject),
	})
}

func (s *DeckStorage) Delete(ctx context.Context, id string) error {
	return s.queries.DeleteDeck(ctx, id)
}

func (s DeckStorage) Search(ctx context.Context, title string) ([]model.Deck, error) {
	sqlcDecks, err := s.queries.SearchDeck(ctx, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	return convertDeckSlice(sqlcDecks)
}

func convertDeck(sqlcDeck sqlc.Deck) (model.Deck, error) {
	modelDeck := model.Deck{
		ID:        sqlcDeck.ID,
		UserID:    sqlcDeck.UserID,
		Title:     sqlcDeck.Title,
		Subject:   sqlcDeck.Subject.String,
		CreatedAt: sqlcDeck.CreatedAt,
		UpdatedAt: sqlcDeck.UpdatedAt,
	}

	if err := json.Unmarshal(sqlcDeck.Cards, &modelDeck.Cards); err != nil {
		return model.Deck{}, err
	}

	return modelDeck, nil
}

func convertDeckSlice(sqlcDecks []sqlc.Deck) ([]model.Deck, error) {
	var modelDecks []model.Deck
	for _, sqlcDeck := range sqlcDecks {
		modelDeck, err := convertDeck(sqlcDeck)
		if err != nil {
			return nil, err
		}

		modelDecks = append(modelDecks, modelDeck)
	}
	return modelDecks, nil
}
