package service

import (
	"context"
	"log"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Session struct {
	storage   *storage.Queries
	validator *validator.Validate
}

func NewSessionService(storage *storage.Queries, validator *validator.Validate) *Session {
	return &Session{
		storage:   storage,
		validator: validator,
	}
}

func (s Session) GetByID(ctx context.Context, id string) (model.Session, error) {
	return s.storage.Session.GetByID(ctx, id)
}

func (s Session) GetByUserID(ctx context.Context, id string) ([]model.Session, error) {
	return s.storage.Session.GetByUserID(ctx, id)
}

func (s Session) Create(ctx context.Context, arg model.SessionCreateParams) (model.Session, error) {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 64)
	if err != nil {
		log.Fatal(err)
	}

	sessionCreateParams := storage.SessionCreateParams{
		ID:          id,
		UserID:      arg.UserID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   arg.ExpiresAt,
	}

	if err = s.storage.Session.Create(ctx, sessionCreateParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s Session) Update(ctx context.Context, id string, arg model.SessionUpdateParams) (model.Session, error) {
	if err := s.storage.Session.Update(ctx, id, arg); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s Session) Delete(ctx context.Context, id string) error {
	return s.storage.Session.Delete(ctx, id)
}
