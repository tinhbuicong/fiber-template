package services

import (
	"context"
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)

// AuthService defines the authentication service interface
type AuthService interface {
	Register(ctx context.Context, req *RegisterInput) (*AuthResult, error)
	Login(ctx context.Context, req *LoginInput) (*AuthResult, error)
	Logout(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*AuthResult, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	VerifyEmail(ctx context.Context, token string) error
	GetProfile(ctx context.Context, userID string) (*UserProfile, error)
	ValidateToken(ctx context.Context, token string) (*TokenClaims, error)
}

// RegisterInput represents registration input
type RegisterInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// LoginInput represents login input
type LoginInput struct {
	Email    string
	Password string
}

// AuthResult represents authentication result
type AuthResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	TokenType    string
	User         *UserProfile
}

// UserProfile represents user profile data
type UserProfile struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Avatar    string
	Role      string
	Verified  bool
	CreatedAt int64
	UpdatedAt int64
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID string
	Email  string
	Role   string
}
