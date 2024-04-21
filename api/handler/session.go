package handler

import (
	"net"
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type SessionHandler struct {
	sessionService model.SessionRepo
}

func NewSessionHandler(sessionService model.SessionRepo) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h SessionHandler) SignUp(c echo.Context) error {
	userCreateParams := new(model.UserCreateParams)
	if err := c.Bind(&userCreateParams); err != nil {
		return err
	}

	authResponse, err := h.sessionService.SignUp(c.Request().Context(), model.SignUpParams{
		Username:    userCreateParams.Username,
		Email:       userCreateParams.Email,
		Password:    userCreateParams.Password,
		Fingerprint: []byte(c.Request().UserAgent()),
		IP:          net.ParseIP(c.RealIP()),
	})
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "invalid credentials"
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

func (h SessionHandler) LogIn(c echo.Context) error {
	email, password, ok := c.Request().BasicAuth()
	if !ok {
		var logInParms presenter.LogInParams
		if err := c.Bind(&logInParms); err != nil {
			return err
		}

		email = logInParms.Email
		password = logInParms.Password
	}

	authResponse, err := h.sessionService.LogIn(c.Request().Context(), model.LogInParams{
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

func (h SessionHandler) GetSession(c echo.Context) error {
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		return err
	}

	session, err := h.sessionService.GetByID(c.Request().Context(), sessionID.Value)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, presenter.Session{
		ID:        session.ID,
		ExpiresAt: session.ExpiresAt,
	})
}
