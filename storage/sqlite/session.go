package sqlite

import (
	"context"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/storage/sqlite/sqlc"
)

type SessionStorage struct {
	queries *sqlc.Queries
}

func NewSessionStorage(db sqlc.DBTX) *SessionStorage {
	return &SessionStorage{queries: sqlc.New(db)}
}

func (s *SessionStorage) GetByID(ctx context.Context, id string) (model.Session, error) {
	sqlcSessions, err := s.queries.GetSessionByID(ctx, id)
	return model.Session(sqlcSessions), err
}

func (s *SessionStorage) GetByUserID(ctx context.Context, userID string) ([]model.Session, error) {
	sqlcSessions, err := s.queries.GetSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	sessions := make([]model.Session, 0, len(sqlcSessions))
	for _, sqlcSession := range sqlcSessions {
		sessions = append(sessions, model.Session(sqlcSession))
	}

	return sessions, nil
}

func (s *SessionStorage) Create(ctx context.Context, arg storage.SessionCreateParams) error {
	return s.queries.CreateSession(ctx, sqlc.CreateSessionParams(arg))
}

func (s *SessionStorage) Update(ctx context.Context, arg storage.SessionUpdateParams) error {
	return s.queries.UpdateSession(ctx, sqlc.UpdateSessionParams(arg))
}

func (s *SessionStorage) Delete(ctx context.Context, id string) error {
	return s.queries.DeleteSession(ctx, id)
}
