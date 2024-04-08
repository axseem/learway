package storage

import (
	"context"

	"github.com/axseem/learway/model"
)

type UserCreateParams struct {
	ID       string
	Username string
	Email    string
	Password []byte
}

type UserUpdatePasswordParams struct {
	Password []byte
	ID       string
}

type UserUpdateUsernameParams struct {
	Username string
	ID       string
}

type UserQueries interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, arg UserCreateParams) error
	UpdatePassword(ctx context.Context, arg UserUpdatePasswordParams) error
	UpdateUsername(ctx context.Context, arg UserUpdateUsernameParams) error
	Delete(ctx context.Context, id string) error
}
