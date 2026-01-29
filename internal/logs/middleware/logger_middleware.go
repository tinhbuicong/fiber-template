package middleware

import (
	logs "fiber-template/internal/logs/models"
	"fiber-template/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ApiLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Proceed to the next handler
		err := c.Next()

		// After the request is finished, capture details
		stop := time.Now()

		logEntry := logs.ApiLog{
			Method:    c.Method(),
			Path:      c.Path(),
			Status:    c.Response().StatusCode(),
			Latency:   stop.Sub(start).String(),
			IP:        c.IP(),
			CreatedAt: stop,
		}

		// Save to Database (Async is better for performance)
		go database.DB.Create(&logEntry)

		return err
	}
}
