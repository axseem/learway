package web

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/handler"
	"github.com/labstack/echo/v4"
)

func attachDB(
	handler func(c echo.Context, db *database.Queries) error,
	db *database.Queries,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler(c, db)
	}
}

func App(db *database.Queries) *echo.Echo {
	e := echo.New()
	e.StaticFS("/assets/", AssetsFS())

	e.GET("/", attachDB(handler.IndexPage, db))
	e.GET("/deck/:id", attachDB(handler.DeckPage, db))
	e.GET("/create", handler.CreatePage)
	e.POST("/create", attachDB(handler.CreateDeck, db))

	return e
}
