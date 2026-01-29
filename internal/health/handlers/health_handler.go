package handlers

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	// 2. Trả về trạng thái khỏe mạnh
	return c.Status(200).JSON(fiber.Map{
		"status": "up",
		"uptime": "active", // Bạn có thể thêm biến tính thời gian chạy của app
	})
}
