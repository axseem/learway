package sqlite

import (
	"context"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/storage/sqlite/sqlc"
)

type UserStorage struct {
	queries *sqlc.Queries
}

func NewUserStorage(db sqlc.DBTX) *UserStorage {
	return &UserStorage{queries: sqlc.New(db)}
}

func (s UserStorage) GetByID(ctx context.Context, id string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByID(ctx, id)
	return model.User(sqlcUser), err
}

func (s UserStorage) GetByEmail(ctx context.Context, email string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByEmail(ctx, email)
	return model.User(sqlcUser), err
}

func (s UserStorage) GetByUsername(ctx context.Context, username string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByUsername(ctx, username)
	return model.User(sqlcUser), err
}

func (s *UserStorage) Create(ctx context.Context, arg storage.UserCreateParams) error {
	return s.queries.CreateUser(ctx, sqlc.CreateUserParams(arg))
}

func (s *UserStorage) UpdatePassword(ctx context.Context, arg storage.UserUpdatePasswordParams) error {
	return s.queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams(arg))
}

func (s *UserStorage) UpdateUsername(ctx context.Context, arg storage.UserUpdateUsernameParams) error {
	return s.queries.UpdateUserUsername(ctx, sqlc.UpdateUserUsernameParams(arg))
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	return s.queries.DeleteUser(ctx, id)
}
