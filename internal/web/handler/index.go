package handler

import (
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) IndexPage(c echo.Context) error {
	rawDecks, err := h.DeckService.List(c.Request().Context())
	if err != nil {
		return err
	}

	return render(c, view.IndexPage(rawDecks))
}
