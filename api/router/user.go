package router

import (
	"github.com/axseem/learway/api/handler"
	"github.com/labstack/echo/v4"
)

func User(api *echo.Group, h *handler.UserHandler) {
	api.GET("/user/:username", h.Get)
	api.GET("/user/:username/decks", h.GetDecks)
}
