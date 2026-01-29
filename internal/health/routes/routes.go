package routes

import (
	"fiber-template/internal/health/handlers"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(router fiber.Router) {
	router.Get("/", handlers.HealthCheck)
}
