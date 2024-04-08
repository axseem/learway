package memory

import (
	"context"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
)

type SessionStorage []model.Session

func (s SessionStorage) GetByID(ctx context.Context, id string) (model.Session, error) {
	for _, session := range s {
		if session.ID == id {
			return session, nil
		}
	}

	return model.Session{}, ErrNotFound
}

func (s SessionStorage) GetByUserID(ctx context.Context, userID string) ([]model.Session, error) {
	var sessions []model.Session

	for _, session := range s {
		if session.UserID == userID {
			sessions = append(sessions, session)
		}
	}

	if len(sessions) == 0 {
		return nil, ErrNotFound
	}

	return sessions, nil
}

func (s *SessionStorage) Create(ctx context.Context, arg storage.SessionCreateParams) error {
	newSession := model.Session{
		ID:          arg.ID,
		UserID:      arg.UserID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   arg.ExpiresAt,
		CreatedAt:   time.Now(),
	}

	*s = append(*s, newSession)

	return nil
}

func (s *SessionStorage) Update(ctx context.Context, arg storage.SessionUpdateParams) error {
	for i, session := range *s {
		if session.ID == arg.ID {
			(*s)[i].Fingerprint = arg.Fingerprint
			(*s)[i].IP = arg.IP
			(*s)[i].ExpiresAt = arg.ExpiresAt
			return nil
		}
	}

	return ErrNotFound
}

func (s *SessionStorage) Delete(ctx context.Context, id string) error {
	for i, session := range *s {
		if session.ID == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}
