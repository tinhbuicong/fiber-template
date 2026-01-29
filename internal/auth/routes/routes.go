package routes

import (
	"fiber-template/internal/auth/handlers"
	"fiber-template/internal/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

// MapRoutes configures auth routes on the given router (caller passes the /auth group with limiter).
// Paths: /register, /login, ... => full path /api/v1/auth/register, /api/v1/auth/login, ...
func MapRoutes(router fiber.Router) {
	authHandler := handlers.NewAuthHandler()
	authMiddleware := middleware.NewAuthMiddleware()

	// Public routes (no authentication required)
	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)
	router.Post("/forgot-password", authHandler.ForgotPassword)
	router.Post("/reset-password", authHandler.ResetPassword)
	router.Get("/verify-email", authHandler.VerifyEmail)
	router.Post("/refresh", authHandler.RefreshToken)

	// Protected routes (authentication required)
	protected := router.Group("", authMiddleware.Authenticate())
	protected.Post("/logout", authHandler.Logout)
	protected.Get("/profile", authHandler.GetProfile)
}
