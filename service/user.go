package service

import (
	"context"
	"log"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/storage"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UserService struct {
	storage *storage.Queries
}

func NewUserService(storage *storage.Queries) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s UserService) GetByID(ctx context.Context, id string) (model.User, error) {
	return s.storage.User.GetByID(ctx, id)
}

func (s UserService) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return s.storage.User.GetByUsername(ctx, username)
}

func (s UserService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return s.storage.User.GetByEmail(ctx, email)
}

func (s UserService) Create(ctx context.Context, arg model.UserCreateParams) (model.User, error) {
	id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		log.Fatal(err)
	}

	passwordHash, err := security.HashPassword(arg.Password)
	if err != nil {
		return model.User{}, err
	}

	userCreateParams := storage.UserCreateParams{
		ID:       id,
		Username: arg.Username,
		Email:    arg.Email,
		Password: passwordHash,
	}

	if err = s.storage.User.Create(ctx, userCreateParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, id)
}

func (s UserService) UpdatePassword(ctx context.Context, arg model.UserUpdatePasswordParams) (model.User, error) {
	rawUser, err := s.storage.User.GetByID(ctx, arg.ID)
	if err != nil {
		return model.User{}, err
	}

	if err = security.CompareHashAndPassword(arg.CurrentPassword, rawUser.Password); err != nil {
		return model.User{}, err
	}

	newPasswordHash, err := security.HashPassword(arg.NewPassword)
	if err != nil {
		return model.User{}, err
	}

	userUpdatePasswordParams := storage.UserUpdatePasswordParams{
		Password: newPasswordHash,
		ID:       rawUser.ID,
	}

	if err = s.storage.User.UpdatePassword(ctx, userUpdatePasswordParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, rawUser.ID)
}

func (s UserService) UpdateUsername(ctx context.Context, arg model.UserUpdateUsernameParams) (model.User, error) {
	userUpdateUsernameParams := storage.UserUpdateUsernameParams{
		Username: arg.Username,
		ID:       arg.ID,
	}

	if err := s.storage.User.UpdateUsername(ctx, userUpdateUsernameParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, arg.ID)
}

func (s UserService) Delete(ctx context.Context, id string) error {
	return s.storage.Deck.Delete(ctx, id)
}
