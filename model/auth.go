package model

import (
	"context"
	"net"
)

type SignUpParams struct {
	Username    string `json:"username" validate:"required,alphanum,gt=0,lte=64"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Fingerprint []byte `json:"fingerprint"`
	IP          net.IP `json:"ip"`
}

type LogInParams struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Fingerprint []byte `json:"fingerprint"`
	IP          net.IP `json:"ip"`
}

type AuthReturnValues struct {
	Session  Session `json:"session"`
	Username string  `json:"username"`
}

type AuthRepo interface {
	SignUp(ctx context.Context, arg SignUpParams) (AuthReturnValues, error)
	LogIn(ctx context.Context, arg LogInParams) (AuthReturnValues, error)
	GetSessionData(ctx context.Context, sessionID string) (Session, error)
}
