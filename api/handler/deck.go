package handler

import (
	"net/http"

	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type DeckHandler struct {
	DeckService model.DeckRepo
}

func NewDeckHandler(deckService model.DeckRepo) *DeckHandler {
	return &DeckHandler{
		DeckService: deckService,
	}
}

func (h DeckHandler) List(c echo.Context) error {
	decks, err := h.DeckService.List(c.Request().Context())
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, decks)
	return nil
}

func (h DeckHandler) Get(c echo.Context) error {
	id := c.Param("id")

	deck, err := h.DeckService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, deck)
	return nil
}

func (h DeckHandler) Create(c echo.Context) error {
	var arg model.DeckCreateParams
	if err := c.Bind(&arg); err != nil {
		return err
	}

	deck, err := h.DeckService.Create(c.Request().Context(), arg)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to create deck"
		return echo.NewHTTPError(code, message)
	}

	c.JSON(http.StatusCreated, deck)
	return nil
}

func (h DeckHandler) Update(c echo.Context) error {
	id := c.Param("id")

	var arg model.DeckCreateParams
	if err := c.Bind(&arg); err != nil {
		return err
	}

	deck, err := h.DeckService.Update(c.Request().Context(), id, arg)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to update deck"
		return echo.NewHTTPError(code, message)
	}

	c.JSON(http.StatusCreated, deck)
	return nil
}

func (h DeckHandler) Delete(c echo.Context) error {
	decks, err := h.DeckService.List(c.Request().Context())
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to list decks"
		return echo.NewHTTPError(code, message)
	}

	c.JSON(http.StatusOK, decks)
	return nil
}
