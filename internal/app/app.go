package app

import (
	"github.com/gui-marc/go-fast-chat/config"
	"github.com/gui-marc/go-fast-chat/internal/handler"
	"github.com/gui-marc/go-fast-chat/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(cfg *config.Config) *echo.Echo {
	e := echo.New()

	// Global Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Dependencies
	db := repository.InitDatabase(cfg.DatabaseURL)
	redis := repository.InitRedis(cfg.RedisURL)

	// Pass dependencies to handlers
	h := handler.NewHandler(db, redis)

	// Register all routes
	handler.RegisterRoutes(e, h)

	return e
}
