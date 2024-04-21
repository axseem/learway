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

type SignUpParams struct {
	Username    string `json:"username" validate:"required,alphanum,gt=0,lte=64"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Fingerprint []byte `json:"fingerprint"`
	IP          net.IP `json:"ip"`
}

type LogInParams struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Fingerprint []byte `json:"fingerprint"`
	IP          net.IP `json:"ip"`
}

type AuthReturnValues struct {
	Session  Session `json:"session"`
	Username string  `json:"username"`
}

type SessionRepo interface {
	GetByID(ctx context.Context, id string) (Session, error)
	GetByUserID(ctx context.Context, id string) ([]Session, error)
	Create(ctx context.Context, arg SessionCreateParams) (Session, error)
	Update(ctx context.Context, id string, arg SessionUpdateParams) (Session, error)
	Delete(ctx context.Context, id string) error
	SignUp(ctx context.Context, arg SignUpParams) (AuthReturnValues, error)
	LogIn(ctx context.Context, arg LogInParams) (AuthReturnValues, error)
}
