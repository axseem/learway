package handler

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func IndexPage(db *database.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		rawDecks, err := db.ListDecks(c.Request().Context())
		if err != nil {
			return err
		}
		return render(c, view.IndexPage(rawDecks))
	}
}
