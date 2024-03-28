package api

import (
	"github.com/axseem/learway/internal/api/handler"
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/service"
	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo, db *database.Queries) {
	g := e.Group("/api")

	h := handler.NewBaseHandler(*service.NewDeckService(db))

	g.POST("/create", h.CreateDeck)
}
