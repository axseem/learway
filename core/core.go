// Package core provides App data structure.
package core

import (
	"github.com/axseem/learway/service"
	"github.com/axseem/learway/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo      *echo.Echo
	Queries   *storage.Queries
	Validator *validator.Validate
	Service   *service.Base
}

func NewApp(q *storage.Queries) *App {
	e := echo.New()
	v := validator.New(validator.WithRequiredStructEnabled())
	s := service.NewBase(q, v)

	return &App{
		Echo:      e,
		Queries:   q,
		Validator: v,
		Service:   s,
	}
}

func (app App) Serve(address string) error {
	return app.Echo.Start(address)
}
