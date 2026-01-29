package routes

import (
	"fiber-template/internal/logs/handlers"
	"fiber-template/internal/logs/repository/postgres"
	"fiber-template/internal/logs/services"
	"fiber-template/pkg/database"

	"github.com/gofiber/fiber/v2"
)

// MapRoutes registers log routes on the given router (e.g. api.Group("/api/v1"))
func MapRoutes(router fiber.Router) {
	repo := postgres.NewLogRepository(database.DB)
	svc := services.NewLogService(repo)
	handler := handlers.NewLogHandler(svc)

	logs := router.Group("/logs")
	logs.Get("/", handler.GetLogs)
}
