package middleware

import (
	"fiber-template/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiterWith returns a rate limiter for a given max and expiration (per-zone use).
func RateLimiterWith(max int, expiration time.Duration) fiber.Handler {
	if max < 1 {
		max = 60
	}
	if expiration < time.Second {
		expiration = time.Minute
	}
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Too Many Requests",
				"message": "Rate limit exceeded. Please try again later.",
			})
		},
	})
}

// RateLimiter returns Fiber rate limiter middleware (legacy single-zone).
// Cấu hình qua env: RATELIMIT_MAX, RATELIMIT_EXPIRATION, RATELIMIT_ENABLED.
func RateLimiter(cfg *config.RateLimitConfig) fiber.Handler {
	if cfg == nil || !cfg.Enabled {
		return func(c *fiber.Ctx) error { return c.Next() }
	}
	max := cfg.Max
	if max < 1 {
		max = 60
	}
	expiration := cfg.Expiration
	if expiration < time.Second {
		expiration = time.Minute
	}
	return RateLimiterWith(max, expiration)
}
