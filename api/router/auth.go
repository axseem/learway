package router

import (
	"github.com/axseem/learway/api/handler"
	"github.com/labstack/echo/v4"
)

func Auth(api *echo.Group, h *handler.AuthHandler) {
	api.POST("/signup", h.SignUp)
	api.POST("/login", h.LogIn)
	api.GET("/session", h.GetSession)
}
