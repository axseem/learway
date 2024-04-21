package handler

import (
	"log"
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService model.UserRepo
	deckService model.DeckRepo
}

func NewUserHandler(userService model.UserRepo, deckService model.DeckRepo) *UserHandler {
	return &UserHandler{
		userService: userService,
		deckService: deckService,
	}
}

func (h UserHandler) Get(c echo.Context) error {
	username := c.Param("username")

	user, err := h.userService.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, presenter.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
}

func (h UserHandler) GetDecks(c echo.Context) error {
	username := c.Param("username")

	user, err := h.userService.GetByUsername(c.Request().Context(), username)
	if err != nil {
		log.Println(err)
		return err
	}

	decks, err := h.deckService.GetByUserID(c.Request().Context(), user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, decks)
}
