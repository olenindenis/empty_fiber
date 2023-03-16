package handlers

import (
	"net/http"

	"envs/internal/core/ports"
	"envs/pkg/cache"
	"envs/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type HealthChecksHandlers struct {
	cache        *cache.Cache
	dbConnection database.Connection
}

var _ ports.HealthChecksHandlers = (*HealthChecksHandlers)(nil)

func NewHealthChecksHandlers(cache *cache.Cache, dbConnection database.Connection) *HealthChecksHandlers {
	return &HealthChecksHandlers{
		cache:        cache,
		dbConnection: dbConnection,
	}
}

// HealthChecks godoc
// @Tags healthChecks
// @Summary healthChecks
// @Description healthChecks
// @Accept  json
// @Produce  json
// @Success 200 {object} HTTPSuccess "ok"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /health_checks [get]
func (h *HealthChecksHandlers) HealthChecks(ctx *fiber.Ctx) error {
	var cacheStatus, dbStatus string
	cacheStatus = http.StatusText(http.StatusOK)
	dbStatus = http.StatusText(http.StatusOK)

	err := h.cache.Ping()
	if err != nil {
		cacheStatus = err.Error()
	}

	conn, err := h.dbConnection.Connection()
	if err != nil {
		dbStatus = err.Error()
	}
	err = conn.Ping()
	if err != nil {
		dbStatus = err.Error()
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"cache":    cacheStatus,
		"database": dbStatus,
		"app":      http.StatusText(http.StatusOK),
	})
}
