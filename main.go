package main

import (
	"fiber-template/config"
	logMiddleware "fiber-template/internal/logs/middleware"
	"fiber-template/internal/logs/models"
	"fiber-template/internal/logs/repository/postgres"
	"fiber-template/internal/logs/services"
	"fiber-template/internal/middleware"
	"fiber-template/internal/routes"
	"fiber-template/pkg/database"
	"fiber-template/pkg/logger"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func errString(err error) string {
	if err == nil {
		return "unknown error"
	}
	return err.Error()
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.Load()

	// Database
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Database: %v", err)
	}
	if database.DB == nil {
		log.Fatal("database.DB is nil after ConnectDB")
	}
	if err := database.DB.AutoMigrate(&models.ApiLog{}); err != nil {
		log.Fatalf("AutoMigrate api_logs: %v", err)
	}
	log.Println("Database connection established and Migrated")

	// Zap + Lumberjack: daily rotation, 7 days, compress, ./storage/logs/app.log
	if err := logger.Init(&cfg.Log); err != nil {
		log.Fatalf("Init logger: %v", err)
	}
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if c == nil {
				log.Printf("[ErrorHandler] nil context, err=%v", err)
				return nil
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errString(err),
			})
		},
	})

	// Log service for middleware (async save to DB + file log)
	logRepo := postgres.NewLogRepository(database.DB)
	logSvc := services.NewLogService(logRepo)

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(cors.New())
	app.Use(middleware.RateLimiter(&cfg.RateLimit))
	app.Use(logMiddleware.ApiLogger(logSvc))

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:" + port))
}
