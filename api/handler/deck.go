package handler

import (
	"net/http"

	"github.com/axseem/learway/api/presenter"
	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

type Deck struct {
	service model.DeckRepo
}

func NewDeck(service model.DeckRepo) *Deck {
	return &Deck{
		service: service,
	}
}

func (h Deck) List(c echo.Context) error {
	decks, err := h.service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, decks)
}

func (h Deck) Get(c echo.Context) error {
	id := c.Param("id")

	deck, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, deck)
}

func (h Deck) Create(c echo.Context) error {
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		return err
	}

	var arg presenter.DeckCreateParams
	if err := c.Bind(&arg); err != nil {
		return err
	}

	deck, err := h.service.Create(c.Request().Context(), model.DeckCreateParams{
		SessionID: sessionID.Value,
		Title:     arg.Title,
		Cards:     arg.Cards,
	})
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to create deck"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusCreated, deck)
}

func (h Deck) Update(c echo.Context) error {
	id := c.Param("id")

	var arg model.DeckUpdateParams
	if err := c.Bind(&arg); err != nil {
		return err
	}

	deck, err := h.service.Update(c.Request().Context(), id, arg)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to update deck"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusCreated, deck)
}

func (h Deck) Delete(c echo.Context) error {
	decks, err := h.service.List(c.Request().Context())
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to list decks"
		return echo.NewHTTPError(code, message)
	}

	return c.JSON(http.StatusOK, decks)
}

func (h Deck) Search(c echo.Context) error {
	title := c.Param("title")

	decks, err := h.service.Search(c.Request().Context(), title)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, decks)
}

func (h Deck) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	decks, err := h.service.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, decks)
}
