package middleware

import (
	"fiber-template/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ApiLogger ghi log request ra file JSON (Zap + Lumberjack) để Promtail/Loki thu thập.
// Không lưu log vào Postgres; Postgres chỉ dùng cho User, Auth và dữ liệu nghiệp vụ.
func ApiLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latencyMs := time.Since(start).Milliseconds()
		status := 0
		method, path, ip := "", "", ""
		if c != nil {
			method, path, ip = c.Method(), c.Path(), c.IP()
			if c.Response() != nil {
				status = c.Response().StatusCode()
			}
		}

		// Ghi file JSON (Loki parse field tự động) – recover để tránh panic từ Zap
		func() {
			defer func() {
				if r := recover(); r != nil {
					_ = r
				}
			}()
			fields := []zap.Field{
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.Int64("latency_ms", latencyMs),
				zap.String("ip", ip),
			}
			if err != nil {
				fields = append(fields, zap.Error(err))
				logger.Error("request_error", fields...)
			} else {
				logger.Info("request", fields...)
			}
		}()

		return err
	}
}
