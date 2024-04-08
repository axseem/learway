package handler

import (
	"encoding/json"
	"io"
	"net"
	"net/http"

	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) SignUp(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to read body"
		return echo.NewHTTPError(code, message)
	}

	UserCreateParams := model.UserCreateParams{}

	if err = json.Unmarshal(b, &UserCreateParams); err != nil {
		code := echo.ErrBadRequest.Code
		message := "failed to parse body"
		return echo.NewHTTPError(code, message)
	}

	_, err = h.UserService.Create(c.Request().Context(), UserCreateParams)
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "ivalid credentials"
		return echo.NewHTTPError(code, message)
	}

	sessionCreateParams := model.SessionCreateParams{
		Email:       UserCreateParams.Email,
		Password:    UserCreateParams.Password,
		Fingerprint: []byte(c.Request().UserAgent()),
		IP:          net.ParseIP(c.RealIP()),
	}

	session, err := h.SessionService.Create(c.Request().Context(), sessionCreateParams)
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "ivalid credentials"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusOK, session)

}

func (h BaseHandeler) LogIn(c echo.Context) error {
	email, password, ok := c.Request().BasicAuth()
	if !ok {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			code := echo.ErrInternalServerError.Code
			message := "failed to read body"
			return echo.NewHTTPError(code, message)
		}

		var credentials struct {
			Email    string
			Password string
		}

		if err = json.Unmarshal(b, &credentials); err != nil {
			code := echo.ErrBadRequest.Code
			message := "failed to parse body"
			return echo.NewHTTPError(code, message)
		}

		email = credentials.Email
		password = credentials.Password
	}

	sessionCreateParams := model.SessionCreateParams{
		Email:       email,
		Password:    password,
		Fingerprint: []byte(c.Request().UserAgent()),
		IP:          net.ParseIP(c.RealIP()),
	}

	session, err := h.SessionService.Create(c.Request().Context(), sessionCreateParams)
	if err != nil {
		code := echo.ErrBadRequest.Code
		message := "ivalid credentials"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusOK, session)
}
