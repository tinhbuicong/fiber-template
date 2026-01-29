package repository

import (
	"context"
	"time"
)

// TokenType represents the type of token
type TokenType string

const (
	TokenTypeRefresh      TokenType = "refresh"
	TokenTypePasswordReset TokenType = "password_reset"
	TokenTypeEmailVerify  TokenType = "email_verify"
)

// Token represents a stored token
type Token struct {
	ID        string
	UserID    string
	Token     string
	Type      TokenType
	ExpiresAt time.Time
	CreatedAt time.Time
	Used      bool
}

// TokenRepository defines the token repository interface
type TokenRepository interface {
	// Create stores a new token
	Create(ctx context.Context, token *Token) error

	// FindByToken finds a token by its value
	FindByToken(ctx context.Context, tokenValue string, tokenType TokenType) (*Token, error)

	// FindByUserID finds tokens by user ID
	FindByUserID(ctx context.Context, userID string, tokenType TokenType) ([]*Token, error)

	// Delete deletes a token by ID
	Delete(ctx context.Context, id string) error

	// DeleteByUserID deletes all tokens for a user
	DeleteByUserID(ctx context.Context, userID string, tokenType TokenType) error

	// MarkAsUsed marks a token as used
	MarkAsUsed(ctx context.Context, id string) error

	// DeleteExpired deletes all expired tokens
	DeleteExpired(ctx context.Context) error

	// IsBlacklisted checks if a token is blacklisted
	IsBlacklisted(ctx context.Context, tokenValue string) (bool, error)

	// Blacklist adds a token to blacklist
	Blacklist(ctx context.Context, tokenValue string, expiresAt time.Time) error
}
