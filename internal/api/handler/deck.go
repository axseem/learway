package handler

import (
	"encoding/json"
	"io"

	"github.com/axseem/learway/internal/model"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) CreateDeck(c echo.Context) error {
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

	if err = h.DeckService.ValidateCreateParams(arg); err != nil {
		code := echo.ErrBadRequest.Code
		return echo.NewHTTPError(code, err)
	}

	deck, err := h.DeckService.Create(c.Request().Context(), arg)
	if err != nil {
		code := echo.ErrInternalServerError.Code
		message := "failed to create deck"
		return echo.NewHTTPError(code, message)
	}

	c.JSON(200, deck)
	return nil
}
