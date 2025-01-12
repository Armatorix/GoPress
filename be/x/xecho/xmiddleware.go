package xecho

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
)

func ErrorLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			slog.Info(
				"request failed",
				slog.String("error", eris.ToString(err, true)),
			)
		}
		return err
	}
}
