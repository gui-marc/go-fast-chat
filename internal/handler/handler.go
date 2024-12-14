package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	DB    *sql.DB
	Redis *redis.Client
}

func NewHandler(db *sql.DB, redis *redis.Client) *Handler {
	return &Handler{
		DB:    db,
		Redis: redis,
	}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok"})
}
