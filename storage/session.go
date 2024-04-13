package storage

import (
	"context"
	"net"
	"time"

	"github.com/axseem/learway/model"
)

type SessionCreateParams struct {
	ID          string
	UserID      string
	Fingerprint []byte
	IP          net.IP
	ExpiresAt   time.Time
}

type SessionUpdateParams struct {
	Fingerprint []byte
	IP          net.IP
	ExpiresAt   time.Time
}

type SessionRepo interface {
	GetByID(ctx context.Context, id string) (model.Session, error)
	GetByUserID(ctx context.Context, userID string) ([]model.Session, error)
	Create(ctx context.Context, arg SessionCreateParams) error
	Update(ctx context.Context, id string, arg SessionUpdateParams) error
	Delete(ctx context.Context, id string) error
}
