package service

import (
	"context"
	"log"
	"time"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/model"
	"github.com/axseem/learway/internal/security"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type SessionService struct {
	db *database.Queries
}

func NewSessionService(db *database.Queries) *SessionService {
	return &SessionService{
		db: db,
	}
}

func (s SessionService) GetByID(ctx context.Context, id string) (model.Session, error) {
	raw, err := s.db.GetSessionByID(ctx, id)
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{
		ID:          raw.ID,
		UserID:      raw.UserID,
		Fingerprint: raw.Fingerprint,
		IP:          raw.IP,
		ExpiresAt:   raw.ExpiresAt,
		CreatedAt:   raw.CreatedAt,
	}, nil
}

func (s SessionService) GetByUserID(ctx context.Context, id string) ([]model.Session, error) {
	raw, err := s.db.GetSessionsByUserID(ctx, id)
	if err != nil {
		return []model.Session{}, err
	}

	sessions := make([]model.Session, len(raw))
	for _, rawSession := range raw {
		sessions = append(sessions, model.Session{
			ID:          rawSession.ID,
			UserID:      rawSession.UserID,
			Fingerprint: rawSession.Fingerprint,
			IP:          rawSession.IP,
			ExpiresAt:   rawSession.ExpiresAt,
			CreatedAt:   rawSession.CreatedAt,
		})
	}

	return sessions, nil
}

func (s SessionService) Create(ctx context.Context, arg model.SessionCreateParams) (model.Session, error) {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 64)
	if err != nil {
		log.Fatal(err)
	}

	rawUser, err := s.db.GetUserByEmail(ctx, arg.Email)
	if err != nil {
		return model.Session{}, err
	}

	if err = security.CompareHashAndPassword(arg.Password, rawUser.Password); err != nil {
		return model.Session{}, err
	}

	createSessionParams := database.CreateSessionParams{
		ID:          id,
		UserID:      rawUser.ID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
	}

	if err = s.db.CreateSession(ctx, createSessionParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s SessionService) Update(ctx context.Context, arg model.SessionUpdateParams) (model.Session, error) {
	updateSessionParams := database.UpdateSessionParams{
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
		ID:          arg.ID,
	}

	if err := s.db.UpdateSession(ctx, updateSessionParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, arg.ID)
}

func (s SessionService) Delete(ctx context.Context, id string) error {
	return s.db.DeleteSession(ctx, id)
}
