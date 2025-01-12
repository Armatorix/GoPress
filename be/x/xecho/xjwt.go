package xecho

import (
	"strings"

	"github.com/Armatorix/GoPress/be/x/xjwt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(secret),
			Skipper: func(c echo.Context) bool {
				return strings.HasSuffix(c.Request().URL.Path, "/callback")
			},
			NewClaimsFunc: func(echo.Context) jwt.Claims {
				return &xjwt.Claims{}
			},
			SuccessHandler: func(c echo.Context) {
				claims, ok := c.Get("user").(*jwt.Token)
				if !ok {
					panic("invalid user claims")
				}

				appClaims, ok := claims.Claims.(*xjwt.Claims)
				if !ok {
					panic("invalid app claims")
				}
				// user for strict server
				SetBaseContextValue(c, CtxKeyUserId, User{
					Authed:   true,
					AuthedAt: appClaims.IssuedAt.Time,
					ID:       appClaims.UserId,
				})
				c.Set(string(CtxKeyUserId), appClaims.UserId)
			},
		})
}
