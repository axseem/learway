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

func NewUserStorage(queries *sqlc.Queries) *UserStorage {
	return &UserStorage{queries: queries}
}

func (s UserStorage) GetByID(ctx context.Context, id string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByID(ctx, id)
	return convertUser(sqlcUser), err
}

func (s UserStorage) GetByEmail(ctx context.Context, email string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByEmail(ctx, email)
	return convertUser(sqlcUser), err
}

func (s UserStorage) GetByUsername(ctx context.Context, username string) (model.User, error) {
	sqlcUser, err := s.queries.GetUserByUsername(ctx, username)
	return convertUser(sqlcUser), err
}

func (s *UserStorage) Create(ctx context.Context, arg storage.UserCreateParams) error {
	return s.queries.CreateUser(ctx, sqlc.CreateUserParams(arg))
}

func (s *UserStorage) UpdatePassword(ctx context.Context, id string, password []byte) error {
	return s.queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		Password: password,
		ID:       id,
	})
}

func (s *UserStorage) UpdateUsername(ctx context.Context, id, username string) error {
	return s.queries.UpdateUserUsername(ctx, sqlc.UpdateUserUsernameParams{
		Username: username,
		ID:       id,
	})
}

func (s *UserStorage) UpdateProfile(ctx context.Context, id string, arg model.UserUpdateProfileParams) error {
	return s.queries.UpdateUserProfile(ctx, sqlc.UpdateUserProfileParams{
		Name:        arg.Name,
		Description: toNullString(arg.Description),
		Picture:     toNullString(arg.Picture),
		ID:          id,
	})
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	return s.queries.DeleteUser(ctx, id)
}

func convertUser(sqlcUser sqlc.User) model.User {
	return model.User{
		ID:          sqlcUser.ID,
		Username:    sqlcUser.Username,
		Email:       sqlcUser.Email,
		Password:    sqlcUser.Password,
		Name:        sqlcUser.Name,
		Description: sqlcUser.Description.String,
		Picture:     sqlcUser.Picture.String,
	}
}
