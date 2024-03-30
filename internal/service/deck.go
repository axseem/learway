package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

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

func (s DeckService) Create(ctx context.Context, arg model.DeckCreateParams) (*model.Deck, error) {
	id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		return &model.Deck{}, err
	}

	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return &model.Deck{}, err
	}

	err = s.db.CreateDeck(ctx, database.CreateDeckParams{
		ID:    id,
		Title: arg.Title,
		Cards: cardsJSON,
	})
	if err != nil {
		return &model.Deck{}, err
	}

	return s.Get(ctx, id)
}

func (s DeckService) Update(ctx context.Context, id string, arg model.DeckCreateParams) (*model.Deck, error) {
	cardsJSON, err := json.Marshal(arg.Cards)
	if err != nil {
		return &model.Deck{}, err
	}

	err = s.db.UpdateDeck(ctx, database.UpdateDeckParams{
		ID:    id,
		Title: arg.Title,
		Cards: cardsJSON,
	})
	if err != nil {
		return &model.Deck{}, err
	}

	return s.Get(ctx, id)
}

func (s DeckService) Delete(ctx context.Context, id string) error {
	return s.db.DeleteDeck(ctx, id)
}

func (s DeckService) ValidateCreateParams(arg model.DeckCreateParams) error {
	const maxTitleLen = 256
	const maxCardsAmount = 4096
	const maxContentLen = 1024

	if len(arg.Title) == 0 || len(arg.Title) > maxTitleLen {
		return errors.New("title length must be greater than 0 and less than " + strconv.Itoa(maxTitleLen))
	}

	if len(arg.Cards) == 0 || len(arg.Title) > maxCardsAmount {
		return errors.New("cards amount must be greater than 0 and less than " + strconv.Itoa(maxCardsAmount))

	}

	for _, card := range arg.Cards {
		if len(card[0]) <= 0 || len(card[0]) > maxContentLen || len(card[1]) <= 0 || len(card[1]) > maxContentLen {
			return errors.New("cards' content length must be greater than 0 and less than " + strconv.Itoa(maxContentLen))
		}
	}

	return nil
}

func (s DeckService) ValidateUpdateParams(ctx context.Context, arg model.DeckUpdateParams) error {
	_, err := s.Get(ctx, arg.ID)
	if err != nil {
		return err
	}

	return s.ValidateCreateParams(model.DeckCreateParams{
		Title: arg.Title,
		Cards: arg.Cards,
	})
}
