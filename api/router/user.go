package router

import (
	"github.com/axseem/learway/api/handler"
	"github.com/labstack/echo/v4"
)

func User(api *echo.Group, h *handler.User) {
	api.GET("/user/:username", h.Get)
}
