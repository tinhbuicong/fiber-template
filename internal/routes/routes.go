package routes

import (
	authRoutes "fiber-template/internal/auth/routes"
	healthRoutes "fiber-template/internal/health/routes"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	// Create a versioned API group
	api := app.Group("/api/v1")
	// Register sub-routes
	authRoutes.MapRoutes(api)

	defaultRouter := app.Group("/")
	healthRoutes.MapRoutes(defaultRouter)

}
