package router

import (
	"github.com/axseem/learway/api/handler"
	"github.com/labstack/echo/v4"
)

func Deck(api *echo.Group, apiAuth *echo.Group, h *handler.DeckHandler) {
	api.GET("/decks", h.List)
	api.GET("/decks/:title", h.Search)
	api.GET("/deck/:id", h.Get)
	apiAuth.PUT("/deck/:id", h.Update)
	apiAuth.DELETE("/deck/:id", h.Delete)
	apiAuth.POST("/deck", h.Create)
}
