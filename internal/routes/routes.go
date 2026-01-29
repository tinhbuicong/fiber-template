package routes

import (
	"fiber-template/config"
	authRoutes "fiber-template/internal/auth/routes"
	healthRoutes "fiber-template/internal/health/routes"
	"fiber-template/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, rateLimit *config.RateLimitConfig) {
	exp := time.Minute
	noop := func(c *fiber.Ctx) error { return c.Next() }

	// Public (/, /health) – 100 req/min
	var rootLimiter fiber.Handler = noop
	if rateLimit != nil && rateLimit.Enabled && rateLimit.PublicMax > 0 {
		e := rateLimit.PublicExpiration
		if e < time.Second {
			e = exp
		}
		rootLimiter = middleware.RateLimiterWith(rateLimit.PublicMax, e)
	}
	root := app.Group("", rootLimiter)
	healthRoutes.MapRoutes(root)

	api := app.Group("/api/v1")

	// Auth (/api/v1/auth/*) – 5 req/min
	var authLimiter fiber.Handler = noop
	if rateLimit != nil && rateLimit.Enabled && rateLimit.AuthMax > 0 {
		e := rateLimit.AuthExpiration
		if e < time.Second {
			e = exp
		}
		authLimiter = middleware.RateLimiterWith(rateLimit.AuthMax, e)
	}
	authGroup := api.Group("/auth", authLimiter)
	authRoutes.MapRoutes(authGroup)

	// Logs không còn API từ Postgres; xem log trên Grafana (Loki)
}
