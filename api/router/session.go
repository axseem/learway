package router

import (
	"github.com/axseem/learway/api/handler"
	"github.com/labstack/echo/v4"
)

func Session(api *echo.Group, h *handler.SessionHandler) {
	api.POST("/signup", h.SignUp)
	api.POST("/login", h.LogIn)
	api.GET("/session", h.GetSession)
}
