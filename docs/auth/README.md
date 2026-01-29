# Auth Service Documentation

## Overview

The Auth Service provides user authentication and authorization functionality for the Fiber Template application. It implements industry-standard security practices including JWT tokens, password hashing with bcrypt, and email verification.

## Folder Structure

```
fiber-template/
├── config/                          # Application configuration
│   └── config.go                    # Configuration loader (env vars)
│
├── internal/                        # Private application code
│   └── auth/                        # Auth service module
│       ├── dto/                     # Data Transfer Objects
│       │   ├── requests.go          # Request DTOs (LoginRequest, RegisterRequest, etc.)
│       │   └── responses.go         # Response DTOs (AuthResponse, UserResponse, etc.)
│       │
│       ├── handlers/                # HTTP Handlers (Controllers)
│       │   └── auth_handler.go      # Auth endpoints handlers
│       │
│       ├── middleware/              # HTTP Middleware
│       │   └── auth_middleware.go   # JWT authentication middleware
│       │
│       ├── models/                  # Domain Models
│       │   └── user.go              # User entity
│       │
│       ├── repository/              # Data Access Layer
│       │   ├── user_repository.go   # User repository interface
│       │   ├── token_repository.go  # Token repository interface
│       │   └── postgres/            # PostgreSQL implementations
│       │       └── user_repository_impl.go
│       │
│       ├── routes/                  # Route definitions
│       │   └── routes.go            # Auth routes setup
│       │
│       └── services/                # Business Logic Layer
│           ├── auth_service.go      # Auth service interface
│           └── auth_service_impl.go # Auth service implementation
│
├── pkg/                             # Reusable packages (can be imported by external apps)
│   ├── hash/                        # Password hashing utilities
│   │   └── hash.go                  # Bcrypt hasher
│   │
│   ├── jwt/                         # JWT utilities
│   │   └── jwt.go                   # JWT service
│   │
│   └── validator/                   # Input validation
│       └── validator.go             # Struct validator
│
├── docs/                            # Documentation
│   └── auth/
│       ├── README.md                # This file
│       └── API.md                   # API documentation
│
├── go.mod
├── go.sum
└── main.go
```

## Architecture

The auth service follows **Clean Architecture** principles:

```
┌─────────────────────────────────────────────────────────────┐
│                      HTTP Layer                              │
│   (handlers, middleware, routes)                            │
├─────────────────────────────────────────────────────────────┤
│                    Business Layer                            │
│   (services - business logic)                               │
├─────────────────────────────────────────────────────────────┤
│                     Data Layer                               │
│   (repository - database operations)                        │
├─────────────────────────────────────────────────────────────┤
│                    Domain Layer                              │
│   (models, dto - entities and data structures)              │
└─────────────────────────────────────────────────────────────┘
```

### Layer Responsibilities

| Layer | Responsibility |
|-------|----------------|
| **Handlers** | HTTP request/response handling, input validation, calling services |
| **Services** | Business logic, orchestrating operations |
| **Repository** | Database CRUD operations, data persistence |
| **Models** | Domain entities, business rules on entities |
| **DTOs** | Request/response structures, API contracts |
| **Middleware** | Cross-cutting concerns (auth, logging, rate limiting) |

## API Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/v1/auth/register` | Register new user | No |
| POST | `/api/v1/auth/login` | Login user | No |
| POST | `/api/v1/auth/logout` | Logout user | Yes |
| POST | `/api/v1/auth/refresh` | Refresh access token | No |
| POST | `/api/v1/auth/forgot-password` | Request password reset | No |
| POST | `/api/v1/auth/reset-password` | Reset password | No |
| GET | `/api/v1/auth/verify-email` | Verify email address | No |
| GET | `/api/v1/auth/profile` | Get user profile | Yes |

## Configuration

Set these environment variables:

```env
# Server
SERVER_PORT=3000
SERVER_HOST=0.0.0.0

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_db
DB_SSL_MODE=disable

# JWT
JWT_SECRET_KEY=your-super-secret-key-change-in-production
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h
JWT_ISSUER=fiber-template

# Redis (for token blacklisting)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Mail (for email verification)
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=587
MAIL_USER=your_user
MAIL_PASSWORD=your_password
MAIL_FROM=noreply@example.com
MAIL_FROM_NAME=Fiber Template
```

## Security Features

- **Password Hashing**: bcrypt with configurable cost factor
- **JWT Tokens**: Access tokens (short-lived) + Refresh tokens (long-lived)
- **Token Blacklisting**: Redis-based refresh token invalidation
- **Email Verification**: Required before full account access
- **Rate Limiting**: Protect against brute force attacks
- **Input Validation**: Comprehensive request validation

## Usage

### Registering the Auth Routes

```go
package main

import (
    "fiber-template/internal/auth/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    
    // Setup auth routes
    routes.SetupAuthRoutes(app)
    
    app.Listen(":3000")
}
```

### Using Auth Middleware

```go
package main

import (
    "fiber-template/internal/auth/middleware"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    authMiddleware := middleware.NewAuthMiddleware()
    
    // Protected route
    app.Get("/protected", authMiddleware.Authenticate(), func(c *fiber.Ctx) error {
        userID := c.Locals("user_id").(string)
        return c.JSON(fiber.Map{"user_id": userID})
    })
    
    // Admin only route
    app.Get("/admin", 
        authMiddleware.Authenticate(),
        authMiddleware.RequireRole("admin"),
        func(c *fiber.Ctx) error {
            return c.JSON(fiber.Map{"message": "Admin area"})
        },
    )
    
    app.Listen(":3000")
}
```

## Dependencies

```
github.com/gofiber/fiber/v2        # Web framework
github.com/golang-jwt/jwt/v5       # JWT library
github.com/go-playground/validator/v10  # Validation
golang.org/x/crypto/bcrypt         # Password hashing
```

## Testing

```bash
# Run all tests
go test ./internal/auth/... -v

# Run with coverage
go test ./internal/auth/... -cover

# Run specific package tests
go test ./internal/auth/services/... -v
```
