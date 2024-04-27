package service

import (
	"context"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
)

type User struct {
	storage   *storage.Queries
	validator *validator.Validate
}

func NewUserService(storage *storage.Queries, validator *validator.Validate) *User {
	return &User{
		storage:   storage,
		validator: validator,
	}
}

func (s User) GetByID(ctx context.Context, id string) (model.User, error) {
	return s.storage.User.GetByID(ctx, id)
}

func (s User) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return s.storage.User.GetByUsername(ctx, username)
}

func (s User) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return s.storage.User.GetByEmail(ctx, email)
}

func (s User) UpdatePassword(ctx context.Context, id string, arg model.UserUpdatePasswordParams) (model.User, error) {
	if err := security.ValidatePassword(arg.NewPassword); err != nil {
		return model.User{}, err
	}

	user, err := s.storage.User.GetByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	if err = security.CompareHashAndPassword(user.Password, arg.CurrentPassword); err != nil {
		return model.User{}, err
	}

	newPasswordHash, err := security.HashPassword(arg.NewPassword)
	if err != nil {
		return model.User{}, err
	}

	if err = s.storage.User.UpdatePassword(ctx, user.ID, newPasswordHash); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, user.ID)
}

func (s User) UpdateUsername(ctx context.Context, id string, arg model.UsernameUpdateParams) (model.User, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.User{}, err
	}

	if err := s.storage.User.UpdateUsername(ctx, id, arg.Username); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, id)
}

func (s User) UpdateProfile(ctx context.Context, id string, arg model.UserUpdateProfileParams) (model.User, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.User{}, err
	}

	if err := s.storage.User.UpdateProfile(ctx, id, arg); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, id)
}

func (s User) Delete(ctx context.Context, id string) error {
	return s.storage.Deck.Delete(ctx, id)
}
