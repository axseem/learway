package handler

import (
	"encoding/json"
	"io"

	"github.com/axseem/learway/internal/model"
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) CreatePage(c echo.Context) error {
	return render(c, view.CreatePage())
}

// TODO move to api
func (h BaseHandeler) CreateDeck(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	var args model.DeckCreateParams
	if err = json.Unmarshal(b, &args); err != nil {
		return err
	}

	return h.DeckService.Create(c.Request().Context(), args)
}
