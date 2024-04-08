package middleware

import (
	"time"

	"github.com/axseem/learway/storage"
	"github.com/labstack/echo/v4"
)

type SessionMiddlware struct {
	queries storage.SessionQueries
}

func NewSessionMiddlware(queries storage.SessionQueries) *SessionMiddlware {
	return &SessionMiddlware{
		queries: queries,
	}
}

func (s SessionMiddlware) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			code := echo.ErrBadRequest.Code
			message := "no cookie provided"
			return echo.NewHTTPError(code, message)
		}
		session, err := s.queries.GetByID(c.Request().Context(), sessionID.Value)
		if err != nil {
			code := echo.ErrUnauthorized.Code
			message := "non-existent session"
			return echo.NewHTTPError(code, message)
		}

		if session.ExpiresAt.Before(time.Now()) {
			code := echo.ErrUnauthorized.Code
			message := "session has been expired"
			return echo.NewHTTPError(code, message)
		}

		return next(c)
	}
}
