package xecho

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Authed   bool      `json:"authed"`
	AuthedAt time.Time `json:"authed_at"`
	ID       int       `json:"id"`
}

func Base(c echo.Context) (context.Context, User) {
	userId, ok := c.Get(string(CtxKeyUserId)).(int)
	if !ok {
		return c.Request().Context(), User{}
	}
	return c.Request().Context(), User{
		Authed: true,
		ID:     userId,
	}
}

func FetchUser(ctx context.Context) User {
	user, ok := ctx.Value(CtxKeyUserId).(User)
	if !ok {
		return User{}
	}
	return user
}

func SetBaseContextValue(c echo.Context, key CtxKeys, value interface{}) {
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(
				c.Request().Context(),
				key,
				value,
			),
		),
	)
}
