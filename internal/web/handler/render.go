package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, componet templ.Component) error {
	return componet.Render(c.Request().Context(), c.Response())
}
