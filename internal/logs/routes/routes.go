package routes

import (
	"fiber-template/internal/logs/handlers"
	"fiber-template/internal/logs/repository/postgres"
	"fiber-template/internal/logs/services"
	"fiber-template/pkg/database"

	"github.com/gofiber/fiber/v2"
)

// MapRoutes registers log routes on the given router (caller passes the /logs group with limiter).
// Full path: GET /api/v1/logs
func MapRoutes(router fiber.Router) {
	repo := postgres.NewLogRepository(database.DB)
	svc := services.NewLogService(repo)
	handler := handlers.NewLogHandler(svc)

	router.Get("/", handler.GetLogs)
}
