package web

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/handler"
	"github.com/labstack/echo/v4"
)

func App(db *database.Queries) *echo.Echo {
	e := echo.New()
	e.StaticFS("/assets/", AssetsFS())

	e.GET("/", handler.IndexPage(db))
	e.GET("/deck/:id", handler.DeckPage(db))
	e.GET("/create", handler.CreatePage)
	e.POST("/create", handler.CreateDeck(db))

	return e
}
