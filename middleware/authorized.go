package middleware

import (
	"net"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/labstack/echo/v4"
)

func Authorized(sessionStorage storage.SessionRepo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionID, err := c.Cookie("sessionID")
			if err != nil {
				code := echo.ErrBadRequest.Code
				message := "no cookie provided"
				return echo.NewHTTPError(code, message)
			}

			session, err := sessionStorage.GetByID(c.Request().Context(), sessionID.Value)
			if err != nil {
				code := echo.ErrUnauthorized.Code
				message := "non-existent session"
				return echo.NewHTTPError(code, message)
			}

			if session.ExpiresAt.Before(time.Now()) {
				code := echo.ErrForbidden.Code
				message := "session has been expired"
				return echo.NewHTTPError(code, message)
			}

			sessionUpdateParams := model.SessionUpdateParams{
				Fingerprint: []byte(c.Request().UserAgent()),
				IP:          net.IP(c.RealIP()),
				ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
			}

			if err := sessionStorage.Update(c.Request().Context(), session.ID, sessionUpdateParams); err != nil {
				code := echo.ErrInternalServerError.Code
				message := "failed to update session"
				return echo.NewHTTPError(code, message)
			}

			return next(c)
		}
	}
}
