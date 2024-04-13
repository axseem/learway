package handler

import (
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService model.UserRepo
}

func NewUserHandler(userService model.UserRepo) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h UserHandler) Get(c echo.Context) error {
	username := c.Param("username")

	user, err := h.UserService.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, presenter.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
	return nil
}
