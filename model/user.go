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
	Username string `json:"username" validate:"required,alphanum,gt=0,lte=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type UserUpdatePasswordParams struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type UsernameUpdateParams struct {
	Username string `json:"username" validate:"required,alphanum,gt=0,lte=64"`
}

type UserRepo interface {
	GetByID(ctx context.Context, id string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, arg UserCreateParams) (User, error)
	UpdatePassword(ctx context.Context, id string, arg UserUpdatePasswordParams) (User, error)
	UpdateUsername(ctx context.Context, id string, arg UsernameUpdateParams) (User, error)
	Delete(ctx context.Context, id string) error
}
