package service

import (
	"context"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/go-playground/validator/v10"
)

type Auth struct {
	userService    *User
	sessionService *Session
	validator      *validator.Validate
}

func NewAuthService(userService *User, sessionService *Session, validator *validator.Validate) *Auth {
	return &Auth{
		userService:    userService,
		sessionService: sessionService,
		validator:      validator,
	}
}

func (s Auth) SignUp(ctx context.Context, arg model.SignUpParams) (model.AuthReturnValues, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.AuthReturnValues{}, err
	}

	user, err := s.userService.Create(ctx, model.UserCreateParams{
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
	})
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	session, err := s.sessionService.Create(ctx, model.SessionCreateParams{
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

func (s Auth) LogIn(ctx context.Context, arg model.LogInParams) (model.AuthReturnValues, error) {
	if err := s.validator.Struct(arg); err != nil {
		return model.AuthReturnValues{}, err
	}

	user, err := s.userService.GetByEmail(ctx, arg.Email)
	if err != nil {
		return model.AuthReturnValues{}, err
	}

	if err = security.CompareHashAndPassword(arg.Password, user.Password); err != nil {
		return model.AuthReturnValues{}, err
	}

	session, err := s.sessionService.Create(ctx, model.SessionCreateParams{
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

func (s Auth) GetSessionData(ctx context.Context, sessionID string) (model.Session, error) {
	return s.sessionService.GetByID(ctx, sessionID)
}
