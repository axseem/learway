package model

import (
	"context"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    []byte    `json:"password"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserUpdatePasswordParams struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type UserUpdateProfileParams struct {
	Name        string `json:"name" validate:"gt=0,lte=64"`
	Description string `json:"description" validate:"gt=0,lte=256"`
	Picture     string `json:"picture" validate:"url"`
}

type UsernameUpdateParams struct {
	Username string `json:"username" validate:"required,alphanum,gt=0,lte=64"`
}

type UserRepo interface {
	GetByID(ctx context.Context, id string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	UpdatePassword(ctx context.Context, id string, arg UserUpdatePasswordParams) (User, error)
	UpdateUsername(ctx context.Context, id string, arg UsernameUpdateParams) (User, error)
	UpdateProfile(ctx context.Context, id string, arg UserUpdateProfileParams) (User, error)
	Delete(ctx context.Context, id string) error
}
