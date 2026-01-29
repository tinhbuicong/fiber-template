package middleware

import (
	"context"
	"fiber-template/internal/logs/models"
	"fiber-template/internal/logs/services"
	"fiber-template/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ApiLogger returns a Fiber handler that captures request details, latency, and saves to DB asynchronously
// and writes to Zap/Lumberjack file. Uses a Goroutine for DB write so the response is not delayed.
func ApiLogger(logSvc services.LogService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// 1. Cho request đi qua trước
		err := c.Next()

		// 2. Thu thập dữ liệu ngay sau khi Next() xong (nil-safe: c/Response có thể nil sau khi Next)
		latencyMs := time.Since(start).Milliseconds()
		status := 0
		method, path, ip := "", "", ""
		if c != nil {
			method, path, ip = c.Method(), c.Path(), c.IP()
			if c.Response() != nil {
				status = c.Response().StatusCode()
			}
		}
		logData := models.ApiLog{
			Method:    method,
			Path:      path,
			Status:    status,
			Latency:   latencyMs,
			IP:        ip,
			CreatedAt: time.Now(),
		}

		// 3. Lưu DB bất đồng bộ
		if logSvc != nil {
			go func(data models.ApiLog) {
				// Sử dụng context.Background() là chuẩn vì request đã kết thúc
				_ = logSvc.Save(context.Background(), &data)
			}(logData) // Truyền trực tiếp vào tham số goroutine cho an toàn tuyệt đối
		}

		// 4. Ghi File log (Zap) – dùng wrapper + recover để tránh panic từ Zap encoder
		func() {
			defer func() {
				if r := recover(); r != nil {
					// Zap có thể panic (vd. nil encoder); bỏ qua log, không làm sập request
				}
			}()
			fields := []zap.Field{
				zap.String("method", logData.Method),
				zap.String("path", logData.Path),
				zap.Int("status", logData.Status),
				zap.Int64("latency_ms", logData.Latency),
				zap.String("ip", logData.IP),
			}
			if err != nil {
				fields = append(fields, zap.Error(err))
				logger.Error("request_error", fields...)
			} else {
				logger.Info("request_success", fields...)
			}
		}()

		return err
	}
}
