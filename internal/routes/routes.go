package routes

import (
	authRoutes "fiber-template/internal/auth/routes"
	healthRoutes "fiber-template/internal/health/routes"
	logsRoutes "fiber-template/internal/logs/routes"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Root "/" via health module (empty prefix so Get("/", HealthCheck) => GET /)
	root := app.Group("")
	healthRoutes.MapRoutes(root)

	api := app.Group("/api/v1")
	authRoutes.MapRoutes(api)
	logsRoutes.MapRoutes(api)
}
