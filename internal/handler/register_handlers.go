package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.GET("/health", h.HealthCheck)
}
