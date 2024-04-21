// Package api implements API group for echo app.
package api

import (
	"github.com/axseem/learway/api/handler"
	"github.com/axseem/learway/api/router"
	"github.com/axseem/learway/middleware"
	"github.com/axseem/learway/service"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo, queries *storage.Queries) {
	api := e.Group("/api")
	apiAuth := api.Group("", middleware.Authorized(queries.Session))

	v := validator.New(validator.WithRequiredStructEnabled())

	deckService := service.NewDeckService(queries, v)
	userService := service.NewUserService(queries, v)
	sessionService := service.NewSessionService(queries, userService, v)

	deckHandler := handler.NewDeckHandler(deckService)
	userHandler := handler.NewUserHandler(userService, deckService)
	sessionHandler := handler.NewSessionHandler(sessionService)

	router.Deck(api, apiAuth, deckHandler)
	router.User(api, userHandler)
	router.Session(api, sessionHandler)
}
