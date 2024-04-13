package model

import (
	"context"
	"net"
	"time"
)

type Session struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userID"`
	Fingerprint []byte    `json:"fingerprint"`
	IP          net.IP    `json:"ip"`
	ExpiresAt   time.Time `json:"expiresAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SessionCreateParams struct {
	UserID      string    `json:"userID"`
	Fingerprint []byte    `json:"fingerprint"`
	IP          net.IP    `json:"ip"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type SessionUpdateParams struct {
	Fingerprint []byte    `json:"fingerprint"`
	IP          net.IP    `json:"ip"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type SessionRepo interface {
	GetByID(ctx context.Context, id string) (Session, error)
	GetByUserID(ctx context.Context, id string) ([]Session, error)
	Create(ctx context.Context, arg SessionCreateParams) (Session, error)
	Update(ctx context.Context, id string, arg SessionUpdateParams) (Session, error)
	Delete(ctx context.Context, id string) error
}
