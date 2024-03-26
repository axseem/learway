package handler

import (
	"encoding/json"

	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func (h *BaseHandeler) DeckPage(c echo.Context) error {
	id := c.Param("id")

	deck, err := h.DeckService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	cardsJSON, err := json.Marshal(deck.Cards)
	if err != nil {
		return err
	}

	return render(c, view.DeckPage(string(cardsJSON), deck.Title))
}
