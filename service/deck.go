// Package service implements app logic.
package service

import (
	"context"
	"log"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Deck struct {
	storage   *storage.Queries
	validator *validator.Validate
}

func NewDeckService(storage *storage.Queries, validator *validator.Validate) *Deck {
	return &Deck{
		storage:   storage,
		validator: validator,
	}
}

func (s Deck) Get(ctx context.Context, id string) (model.Deck, error) {
	return s.storage.Deck.Get(ctx, id)
}

func (s Deck) GetByUsername(ctx context.Context, username string) ([]model.Deck, error) {
	user, err := s.storage.User.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return s.storage.Deck.GetByUserID(ctx, user.ID)
}

func (s Deck) List(ctx context.Context) ([]model.Deck, error) {
	return s.storage.Deck.List(ctx)
}

func (s Deck) Create(ctx context.Context, arg model.DeckCreateParams) (model.Deck, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.Deck{}, err
	}

	id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		log.Fatal(err)
	}

	session, err := s.storage.Session.GetByID(ctx, arg.SessionID)
	if err != nil {
		return model.Deck{}, err
	}

	err = s.storage.Deck.Create(ctx, storage.DeckCreateParams{
		ID:     id,
		UserID: session.UserID,
		Title:  arg.Title,
		Cards:  arg.Cards,
	})
	if err != nil {
		return model.Deck{}, err
	}

	return s.storage.Deck.Get(ctx, id)
}

func (s Deck) Update(ctx context.Context, id string, arg model.DeckUpdateParams) (model.Deck, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.Deck{}, err
	}

	err := s.storage.Deck.Update(ctx, id, storage.DeckUpdateParams(arg))
	if err != nil {
		return model.Deck{}, err
	}

	return s.Get(ctx, id)
}

func (s Deck) Delete(ctx context.Context, id string) error {
	return s.storage.Deck.Delete(ctx, id)
}

func (s Deck) Search(ctx context.Context, title string) ([]model.Deck, error) {
	return s.storage.Deck.Search(ctx, title)
}
