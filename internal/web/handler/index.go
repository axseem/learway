package handler

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context, db *database.Queries) error {
	rawDecks, err := db.ListDecks(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.IndexPage(rawDecks))
}
