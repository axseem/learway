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
		return err
	}

	var args model.DeckCreateParams
	if err = json.Unmarshal(b, &args); err != nil {
		return err
	}

	if len(args.Title) == 0 || len(args.Title) > 255 {
		message := "title length must be greater than 0 and less than 256"
		return echo.NewHTTPError(echo.ErrBadRequest.Code, message)
	}

	if len(args.Cards) == 0 || len(args.Title) > 4096 {
		message := "cards amount must be greater than 0 and less than 4096"
		return echo.NewHTTPError(echo.ErrBadRequest.Code, message)
	}

	for _, card := range args.Cards {
		if len(card[0]) <= 0 || len(card[0]) > 1024 || len(card[1]) <= 0 || len(card[1]) > 1024 {
			message := "cards' content length must be greater than 0 and less than 1024"
			return echo.NewHTTPError(echo.ErrBadRequest.Code, message)
		}
	}

	deck, err := h.DeckService.Create(c.Request().Context(), args)
	if err != nil {
		return err
	}

	c.JSON(200, deck)
	return nil
}
