package repository

import (
	"context"

	"fiber-template/internal/auth/models"
)

// UserRepository defines the user repository interface
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *models.User) error

	// FindByID finds a user by ID
	FindByID(ctx context.Context, id string) (*models.User, error)

	// FindByEmail finds a user by email
	FindByEmail(ctx context.Context, email string) (*models.User, error)

	// Update updates a user
	Update(ctx context.Context, user *models.User) error

	// Delete deletes a user by ID
	Delete(ctx context.Context, id string) error

	// ExistsByEmail checks if a user exists by email
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// UpdatePassword updates user password
	UpdatePassword(ctx context.Context, id string, hashedPassword string) error

	// VerifyUser marks a user as verified
	VerifyUser(ctx context.Context, id string) error
}
