package model

import (
	"context"
	"time"
)

type Session struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userID"`
	Fingerprint []byte    `json:"fingerprint"`
	IP          int64     `json:"ip"`
	ExpiresAt   time.Time `json:"expiresAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SessionCreateParams struct {
	UserID      string `json:"userID"`
	Fingerprint []byte `json:"fingerprint"`
	IP          string `json:"ip"`
}

type SessionUpdateParams struct {
	Fingerprint []byte `json:"fingerprint"`
	IP          string `json:"ip"`
}

type SessionRepo interface {
	GetByID(ctx context.Context, id string) (Session, error)
	GetByUserID(ctx context.Context, id string) ([]Session, error)
	Create(ctx context.Context, arg SessionCreateParams) (Session, error)
	Update(ctx context.Context, arg SessionUpdateParams) (Session, error)
	Delete(ctx context.Context, id string) error
}
