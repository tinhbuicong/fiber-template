package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware provides JWT authentication middleware
type AuthMiddleware struct {
	// jwtService jwt.JWTService
}

// NewAuthMiddleware creates a new auth middleware instance
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Authenticate validates JWT token and sets user context
func (m *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
				"code":  "AUTH_MISSING_HEADER",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
				"code":  "AUTH_INVALID_FORMAT",
			})
		}

		token := parts[1]

		// TODO: Validate token using JWT service
		// claims, err := m.jwtService.ValidateToken(token)
		// if err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"error": "Invalid or expired token",
		// 		"code":  "AUTH_INVALID_TOKEN",
		// 	})
		// }

		// Set user info in context
		c.Locals("token", token)
		// c.Locals("user_id", claims.UserID)
		// c.Locals("user_email", claims.Email)
		// c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// RequireRole checks if user has required role
func (m *AuthMiddleware) RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role")
		if userRole == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
				"code":  "AUTH_NOT_AUTHENTICATED",
			})
		}

		role := userRole.(string)
		for _, r := range roles {
			if role == r {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Insufficient permissions",
			"code":  "AUTH_FORBIDDEN",
		})
	}
}

// OptionalAuth validates JWT token if present, but doesn't require it
func (m *AuthMiddleware) OptionalAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Next()
		}

		token := parts[1]

		// TODO: Validate token and set context if valid
		c.Locals("token", token)

		return c.Next()
	}
}
