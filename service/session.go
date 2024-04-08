package service

import (
	"context"
	"log"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/storage"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type SessionService struct {
	storage *storage.Queries
}

func NewSessionService(storage *storage.Queries) *SessionService {
	return &SessionService{
		storage: storage,
	}
}

func (s SessionService) GetByID(ctx context.Context, id string) (model.Session, error) {
	return s.storage.Session.GetByID(ctx, id)
}

func (s SessionService) GetByUserID(ctx context.Context, id string) ([]model.Session, error) {
	return s.storage.Session.GetByUserID(ctx, id)
}

func (s SessionService) Create(ctx context.Context, arg model.SessionCreateParams) (model.Session, error) {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 64)
	if err != nil {
		log.Fatal(err)
	}

	rawUser, err := s.storage.User.GetByEmail(ctx, arg.Email)
	if err != nil {
		return model.Session{}, err
	}

	if err = security.CompareHashAndPassword(arg.Password, rawUser.Password); err != nil {
		return model.Session{}, err
	}

	sessionCreateParams := storage.SessionCreateParams{
		ID:          id,
		UserID:      rawUser.ID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
	}

	if err = s.storage.Session.Create(ctx, sessionCreateParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s SessionService) Update(ctx context.Context, arg model.SessionUpdateParams) (model.Session, error) {
	sessionUpdateParams := storage.SessionUpdateParams{
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
		ID:          arg.ID,
	}

	if err := s.storage.Session.Update(ctx, sessionUpdateParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, arg.ID)
}

func (s SessionService) Delete(ctx context.Context, id string) error {
	return s.storage.Session.Delete(ctx, id)
}
