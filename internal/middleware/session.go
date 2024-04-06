package middleware

import (
	"bytes"
	"net"
	"time"

	"github.com/axseem/learway/internal/database"
	"github.com/labstack/echo/v4"
)

type SessionSource struct {
	db *database.Queries
}

func NewSessionSource(db *database.Queries) *SessionSource {
	return &SessionSource{
		db: db,
	}
}

func (s SessionSource) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("token")
		if err != nil {
			return err
		}
		session, err := s.db.GetSessionByID(c.Request().Context(), token.Value)
		if err != nil {
			return err
		}

		if session.ExpiresAt.Before(time.Now()) || !bytes.Equal(session.IP, net.ParseIP(c.RealIP())) {
			return err
		}

		return next(c)
	}
}
