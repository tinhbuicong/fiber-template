package main

import (
	loggerMiddleware "fiber-template/internal/logs/middleware"
	"fiber-template/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	// app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(loggerMiddleware.ApiLogger())

	// Setup auth routes
	routes.RegisterRoutes(app)

	app.Listen("127.0.0.1:" + port)
}
