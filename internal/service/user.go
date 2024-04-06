package service

import (
	"context"
	"log"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/model"
	"github.com/axseem/learway/internal/security"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UserService struct {
	db *database.Queries
}

func NewUserService(db *database.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

func (s UserService) GetByID(ctx context.Context, id string) (model.User, error) {
	raw, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return model.User(raw), nil
}

func (s UserService) GetByUsername(ctx context.Context, username string) (model.User, error) {
	raw, err := s.db.GetUserByUsername(ctx, username)
	if err != nil {
		return model.User{}, err
	}

	return model.User(raw), nil
}

func (s UserService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	raw, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return model.User(raw), nil
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

	createUserParams := database.CreateUserParams{
		ID:       id,
		Username: arg.Username,
		Email:    arg.Email,
		Password: passwordHash,
	}

	if err = s.db.CreateUser(ctx, createUserParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, id)
}

func (s UserService) UpdatePassword(ctx context.Context, arg model.UserUpdatePasswordParams) (model.User, error) {
	rawUser, err := s.db.GetUserByID(ctx, arg.ID)
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

	updateUserPasswordParams := database.UpdateUserPasswordParams{
		Password: newPasswordHash,
		ID:       rawUser.ID,
	}

	if err = s.db.UpdateUserPassword(ctx, updateUserPasswordParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, rawUser.ID)
}

func (s UserService) UpdateUsername(ctx context.Context, arg model.UserUpdateUsernameParams) (model.User, error) {
	updateUserUsernameParams := database.UpdateUserUsernameParams{
		Username: arg.Username,
		ID:       arg.ID,
	}

	if err := s.db.UpdateUserUsername(ctx, updateUserUsernameParams); err != nil {
		return model.User{}, err
	}

	return s.GetByID(ctx, arg.ID)
}

func (s UserService) Delete(ctx context.Context, id string) error {
	return s.db.DeleteUser(ctx, id)
}
