package handler

import (
	"net"
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService model.AuthRepo
}

func NewAuthHandler(authService model.AuthRepo) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h AuthHandler) SignUp(c echo.Context) error {
	userCreateParams := new(model.UserCreateParams)
	if err := c.Bind(&userCreateParams); err != nil {
		return err
	}

	authResponse, err := h.authService.SignUp(c.Request().Context(), model.SignUpParams{
		Username:    userCreateParams.Username,
		Email:       userCreateParams.Email,
		Password:    userCreateParams.Password,
		Fingerprint: []byte(c.Request().UserAgent()),
		IP:          net.ParseIP(c.RealIP()),
	})
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "ivalid credentials"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusOK, presenter.AuthResponse{
		Session: presenter.Session{
			ID:        authResponse.Session.ID,
			ExpiresAt: authResponse.Session.ExpiresAt,
		},
		Username: authResponse.Username,
	})
}

func (h AuthHandler) LogIn(c echo.Context) error {
	email, password, ok := c.Request().BasicAuth()
	if !ok {
		var logInParms presenter.LogInParms
		if err := c.Bind(&logInParms); err != nil {
			return err
		}

		email = logInParms.Email
		password = logInParms.Password
	}

	authResponse, err := h.authService.LogIn(c.Request().Context(), model.LogInParams{
		Email:       email,
		Password:    password,
		Fingerprint: []byte(c.Request().UserAgent()),
		IP:          net.ParseIP(c.RealIP()),
	})
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "ivalid credentials"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusOK, presenter.AuthResponse{
		Session: presenter.Session{
			ID:        authResponse.Session.ID,
			ExpiresAt: authResponse.Session.ExpiresAt,
		},
		Username: authResponse.Username,
	})
}
