package web

import (
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/service"
	"github.com/axseem/learway/internal/web/handler"
	"github.com/labstack/echo/v4"
)

func App(e *echo.Echo, db *database.Queries) {
	e.StaticFS("/assets/", AssetsFS())

	h := handler.NewBaseHandler(*service.NewDeckService(db))

	e.GET("/", h.IndexPage)
	e.GET("/deck/:id", h.DeckPage)
	e.GET("/create", h.CreatePage)
}
