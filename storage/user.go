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

type UserRepo interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, arg UserCreateParams) error
	UpdatePassword(ctx context.Context, id string, password []byte) error
	UpdateUsername(ctx context.Context, id, username string) error
	Delete(ctx context.Context, id string) error
}
