package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Armatorix/GoPress/be/api/admin"
	"github.com/Armatorix/GoPress/be/api/auth"
	"github.com/Armatorix/GoPress/be/api/content"
	"github.com/Armatorix/GoPress/be/api/public"
	"github.com/Armatorix/GoPress/be/config"
	"github.com/Armatorix/GoPress/be/db"
	dbext "github.com/Armatorix/GoPress/be/db/ext"
	"github.com/Armatorix/GoPress/be/pkg/openai"
	"github.com/Armatorix/GoPress/be/x/xecho"
	"github.com/Armatorix/GoPress/be/x/xjwt"
	"github.com/Armatorix/GoPress/be/x/xslog"
	"github.com/go-playground/validator/v10"
	"github.com/golang-cz/devslog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type V struct {
	Validator *validator.Validate
}

func (cv *V) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	logger := slog.New(devslog.NewHandler(os.Stdout, nil))
	// optional: set global logger
	slog.SetDefault(logger)

	cfg, err := config.FromEnv()
	if err != nil {
		xslog.FatalErr(err)
	}

	if cfg.Server.RedirectMode {
		e := echo.New()
		e.Any("*", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, c.Request().Host)
		})
		xslog.FatalErr(e.Start(cfg.Server.Address()))
	}

	dbext.Init([]byte(cfg.DB.DataEncryptionKey))
	db, err := db.NewWithMigrate(cfg.DB)
	if err != nil {
		xslog.FatalErr(err)
	}

	// TODO: user data from config
	err = db.InitAdminUser(context.Background(), "admin")
	if err != nil {
		xslog.FatalErr(err)
	}
	e := echo.New()
	e.Validator = &V{validator.New()}
	apiNoAuth := e.Group("/api/v1")

	openaic := openai.New(cfg.OpenAI)
	dbAuthMiddleware := xecho.NewAuthStoreMiddleware(db)
	apiAuth := apiNoAuth.Group("", xecho.AuthMiddleware(cfg.Auth.JwtSecret), dbAuthMiddleware)

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		xecho.ErrorLoggerMiddleware,
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: cfg.Server.CORSOrigins,
				AllowHeaders: []string{
					echo.HeaderAccept,
					echo.HeaderOrigin,
					echo.HeaderContentType,
					echo.HeaderAuthorization,
					echo.HeaderAccessControlAllowOrigin,
				},
				AllowCredentials: true,
			},
		),
	)

	auth.RegisterHandlers(
		apiNoAuth,
		auth.NewHandler(
			db,
			xjwt.New([]byte(cfg.Auth.JwtSecret)),
		),
	)

	admin.
		RegisterHandlers(
			apiAuth,
			admin.NewHandler(db),
		)

	content.
		RegisterHandlers(
			apiAuth,
			content.NewHandler(db, openaic),
		)

	public.RegisterHandlers(
		apiNoAuth,
		public.NewHandler(db),
	)

	fe := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderCacheControl, "no-store, max-age=0")
			return next(c)
		}
	})

	// TODO: check for admin role
	staticPath := "/app/public"
	err = filepath.Walk(staticPath,
		func(path string, _ os.FileInfo, _ error) error {
			routePath := path[len(staticPath):]
			fe.GET(routePath, func(c echo.Context) error {
				return c.File(path)
			})
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fe.Any("*", func(c echo.Context) error {
		return c.File("/app/public/index.html")
	})

	xslog.FatalErr(e.Start(cfg.Server.Address()))
}
