package handler

import (
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type User struct {
	service model.UserRepo
}

func NewUser(service model.UserRepo) *User {
	return &User{
		service: service,
	}
}

func (h User) Get(c echo.Context) error {
	username := c.Param("username")

	user, err := h.service.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, presenter.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
}
