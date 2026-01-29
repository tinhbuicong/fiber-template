package routes

import (
	"fiber-template/internal/auth/handlers"
	"fiber-template/internal/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

// MapRoutes configures all auth-related routes
func MapRoutes(router fiber.Router) {
	authHandler := handlers.NewAuthHandler()
	authMiddleware := middleware.NewAuthMiddleware()

	// Auth routes group
	auth := router.Group("/auth")

	// Public routes (no authentication required)
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/forgot-password", authHandler.ForgotPassword)
	auth.Post("/reset-password", authHandler.ResetPassword)
	auth.Get("/verify-email", authHandler.VerifyEmail)
	auth.Post("/refresh", authHandler.RefreshToken)

	// Protected routes (authentication required)
	protected := auth.Group("", authMiddleware.Authenticate())
	protected.Post("/logout", authHandler.Logout)
	protected.Get("/profile", authHandler.GetProfile)
}
