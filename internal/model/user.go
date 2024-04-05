package model

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserCreateParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdatePasswordParams struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type UserUpdateUsernameParams struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserRepo interface {
	GetByID(ctx context.Context, id string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, arg UserCreateParams) (User, error)
	UpdatePassword(ctx context.Context, arg UserUpdatePasswordParams) (User, error)
	UpdateUsername(ctx context.Context, arg UserUpdateUsernameParams) (User, error)
	Delete(ctx context.Context, id string) error
}
