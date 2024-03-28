package handler

import (
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
)

func (h BaseHandeler) CreatePage(c echo.Context) error {
	return render(c, view.CreatePage())
}
