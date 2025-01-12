package xecho

import (
	"context"

	"github.com/labstack/echo/v4"
)

type AuthStore interface {
	UserExists(ctx context.Context, id int) (bool, error)
}

func NewAuthStoreMiddleware(store AuthStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := FetchUser(c.Request().Context())
			if user.Authed {
				exists, err := store.UserExists(c.Request().Context(), user.ID)
				if err != nil {
					return err
				}
				if !exists {
					return echo.ErrUnauthorized
				}
			} else {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
