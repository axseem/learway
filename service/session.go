package service

import (
	"context"
	"log"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Session struct {
	storage   *storage.Queries
	validator *validator.Validate
}

func NewSessionService(storage *storage.Queries, validator *validator.Validate) *Session {
	return &Session{
		storage:   storage,
		validator: validator,
	}
}

func (s Session) GetByID(ctx context.Context, id string) (model.Session, error) {
	return s.storage.Session.GetByID(ctx, id)
}

func (s Session) GetByUserID(ctx context.Context, id string) ([]model.Session, error) {
	return s.storage.Session.GetByUserID(ctx, id)
}

func (s Session) CreateUser(ctx context.Context, arg model.UserCreateParams) (model.User, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.User{}, err
	}

	if err := security.ValidatePassword(arg.Password); err != nil {
		return model.User{}, err
	}

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

	return s.storage.User.GetByID(ctx, id)
}

func (s Session) Create(ctx context.Context, arg model.SessionCreateParams) (model.Session, error) {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 64)
	if err != nil {
		log.Fatal(err)
	}

	sessionCreateParams := storage.SessionCreateParams{
		ID:          id,
		UserID:      arg.UserID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   arg.ExpiresAt,
	}

	if err = s.storage.Session.Create(ctx, sessionCreateParams); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s Session) Update(ctx context.Context, id string, arg model.SessionUpdateParams) (model.Session, error) {
	if err := s.storage.Session.Update(ctx, id, arg); err != nil {
		return model.Session{}, err
	}

	return s.GetByID(ctx, id)
}

func (s Session) Delete(ctx context.Context, id string) error {
	return s.storage.Session.Delete(ctx, id)
}

func (s Session) SignUp(ctx context.Context, arg model.SignUpParams) (model.AuthReturnValues, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.AuthReturnValues{}, err
	}

	user, err := s.CreateUser(ctx, model.UserCreateParams{
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
	})
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	session, err := s.Create(ctx, model.SessionCreateParams{
		UserID:      user.ID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
	})
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	return model.AuthReturnValues{
		Session:  session,
		Username: user.Username,
	}, nil
}

func (s Session) LogIn(ctx context.Context, arg model.LogInParams) (model.AuthReturnValues, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.AuthReturnValues{}, err
	}

	user, err := s.storage.User.GetByEmail(ctx, arg.Email)
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	if err = security.CompareHashAndPassword(arg.Password, user.Password); err != nil {
		return model.AuthReturnValues{}, err
	}

	session, err := s.Create(ctx, model.SessionCreateParams{
		UserID:      user.ID,
		Fingerprint: arg.Fingerprint,
		IP:          arg.IP,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
	})
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	return model.AuthReturnValues{
		Session:  session,
		Username: user.Username,
	}, nil
}
