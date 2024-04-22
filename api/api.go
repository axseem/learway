// Package api implements API group for echo app.
package api

import (
	"github.com/axseem/learway/api/handler"
	"github.com/axseem/learway/api/router"
	"github.com/axseem/learway/core"
	"github.com/axseem/learway/middleware"
)

func Attach(app *core.App) {
	api := app.Echo.Group("/api")
	apiAuth := app.Echo.Group("", middleware.Authorized(app.Queries.Session))

	h := handler.NewBase(app.Service)

	router.Deck(api, apiAuth, h.Deck)
	router.User(api, h.User)
	router.Session(api, h.Session)
}
