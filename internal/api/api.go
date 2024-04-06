package api

import (
	"github.com/axseem/learway/internal/api/handler"
	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/middleware"
	"github.com/axseem/learway/internal/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BaseValidator struct {
	Validator *validator.Validate
}

func (bv *BaseValidator) Validate(i interface{}) error {
	return bv.Validator.Struct(i)
}

func API(e *echo.Echo, db *database.Queries) {
	g := e.Group("/api")

	e.Validator = &BaseValidator{Validator: validator.New()}

	h := handler.NewBaseHandler(
		service.NewDeckService(db),
		service.NewUserService(db),
		service.NewSessionService(db),
	)

	session := middleware.NewSessionSource(db)

	a := g.Group("")
	a.Use(session.Authorized)

	g.GET("/decks", h.ListDecks)
	g.GET("/deck/:id", h.GetDeck)
	a.PUT("/deck/:id", h.UpdateDeck)
	a.DELETE("/deck/:id", h.DeleteDeck)
	a.POST("/deck", h.CreateDeck)
	g.POST("/signup", h.SignUp)
	g.POST("/login", h.LogIn)
}
