package handler

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func DeckPage(db *database.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		deck, err := db.GetDeck(c.Request().Context(), id)
		if err != nil {
			return err
		}

		return render(c, view.DeckPage(deck))
	}
}
