package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/axseem/learway/model"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) ListDecks(c echo.Context) error {
	decks, err := h.DeckService.List(c.Request().Context())
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, decks)
	return nil
}

func (h BaseHandeler) GetDeck(c echo.Context) error {
	id := c.Param("id")

	deck, err := h.DeckService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, deck)
	return nil
}

func (h BaseHandeler) CreateDeck(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to read body"
		return echo.NewHTTPError(code, message)
	}

	var arg model.DeckCreateParams
	if err = json.Unmarshal(b, &arg); err != nil {

	}

	if err = c.Validate(arg); err != nil {
		code := echo.ErrBadRequest.Code
		return echo.NewHTTPError(code, err)
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

func (h BaseHandeler) UpdateDeck(c echo.Context) error {
	id := c.Param("id")

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to read body"
		return echo.NewHTTPError(code, message)
	}

	var arg model.DeckCreateParams
	if err = json.Unmarshal(b, &arg); err != nil {
		code := echo.ErrBadRequest.Code
		message := "failed to parse body"
		return echo.NewHTTPError(code, message)
	}

	if err = c.Validate(arg); err != nil {
		code := echo.ErrBadRequest.Code
		log.Println(err)
		return echo.NewHTTPError(code, err)
	}

	deck, err := h.DeckService.Update(c.Request().Context(), id, arg)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		log.Println(err)
		message := "failed to update deck"
		return echo.NewHTTPError(code, message)
	}

	c.JSON(http.StatusCreated, deck)
	return nil
}

func (h BaseHandeler) DeleteDeck(c echo.Context) error {
	decks, err := h.DeckService.List(c.Request().Context())
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, decks)
	return nil
}
