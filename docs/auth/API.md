# Auth Service API Documentation

## Base URL

```
http://localhost:3000/api/v1/auth
```

---

## Endpoints

### 1. Register User

Create a new user account.

**Endpoint:** `POST /register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!",
  "confirm_password": "SecurePass123!",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Validation Rules:**
| Field | Rules |
|-------|-------|
| email | Required, valid email format |
| password | Required, min 8 chars, max 72 chars |
| confirm_password | Required, must match password |
| first_name | Required, min 2 chars, max 50 chars |
| last_name | Required, min 2 chars, max 50 chars |

**Success Response (201):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "token_type": "Bearer",
  "expires_in": 900,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "user",
    "verified": false,
    "created_at": 1706000000,
    "updated_at": 1706000000
  }
}
```

**Error Response (400):**
```json
{
  "error": "Validation failed",
  "code": "VALIDATION_ERROR",
  "fields": [
    {
      "field": "email",
      "message": "email must be a valid email address"
    }
  ]
}
```

**Error Response (409):**
```json
{
  "error": "User already exists",
  "code": "USER_EXISTS"
}
```

---

### 2. Login

Authenticate user and get tokens.

**Endpoint:** `POST /login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!"
}
```

**Success Response (200):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "token_type": "Bearer",
  "expires_in": 900,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "user",
    "verified": true,
    "created_at": 1706000000,
    "updated_at": 1706000000
  }
}
```

**Error Response (401):**
```json
{
  "error": "Invalid credentials",
  "code": "INVALID_CREDENTIALS"
}
```

---

### 3. Logout

Invalidate user session/tokens.

**Endpoint:** `POST /logout`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Success Response (200):**
```json
{
  "message": "Logout successful"
}
```

---

### 4. Refresh Token

Get new access token using refresh token.

**Endpoint:** `POST /refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Success Response (200):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "token_type": "Bearer",
  "expires_in": 900
}
```

**Error Response (401):**
```json
{
  "error": "Invalid or expired refresh token",
  "code": "INVALID_REFRESH_TOKEN"
}
```

---

### 5. Forgot Password

Request password reset email.

**Endpoint:** `POST /forgot-password`

**Request Body:**
```json
{
  "email": "user@example.com"
}
```

**Success Response (200):**
```json
{
  "message": "Password reset email sent"
}
```

> **Note:** Always returns success to prevent email enumeration attacks.

---

### 6. Reset Password

Reset password using token from email.

**Endpoint:** `POST /reset-password`

**Request Body:**
```json
{
  "token": "reset-token-from-email",
  "password": "NewSecurePass123!",
  "confirm_password": "NewSecurePass123!"
}
```

**Success Response (200):**
```json
{
  "message": "Password reset successful"
}
```

**Error Response (400):**
```json
{
  "error": "Invalid or expired reset token",
  "code": "INVALID_RESET_TOKEN"
}
```

---

### 7. Verify Email

Verify email address using token.

**Endpoint:** `GET /verify-email?token=<verification_token>`

**Query Parameters:**
| Parameter | Description |
|-----------|-------------|
| token | Email verification token |

**Success Response (200):**
```json
{
  "message": "Email verified successfully"
}
```

**Error Response (400):**
```json
{
  "error": "Invalid or expired verification token",
  "code": "INVALID_VERIFICATION_TOKEN"
}
```

---

### 8. Get Profile

Get authenticated user's profile.

**Endpoint:** `GET /profile`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Success Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "avatar": "https://example.com/avatar.jpg",
  "role": "user",
  "verified": true,
  "created_at": 1706000000,
  "updated_at": 1706000000
}
```

---

## Error Codes

| Code | Description |
|------|-------------|
| `VALIDATION_ERROR` | Request validation failed |
| `USER_EXISTS` | User with email already exists |
| `USER_NOT_FOUND` | User not found |
| `INVALID_CREDENTIALS` | Email or password is incorrect |
| `INVALID_TOKEN` | JWT token is invalid |
| `TOKEN_EXPIRED` | JWT token has expired |
| `INVALID_REFRESH_TOKEN` | Refresh token is invalid or expired |
| `INVALID_RESET_TOKEN` | Password reset token is invalid or expired |
| `INVALID_VERIFICATION_TOKEN` | Email verification token is invalid or expired |
| `AUTH_MISSING_HEADER` | Authorization header is missing |
| `AUTH_INVALID_FORMAT` | Authorization header format is invalid |
| `AUTH_NOT_AUTHENTICATED` | User is not authenticated |
| `AUTH_FORBIDDEN` | User does not have required permissions |

---

## Authentication

Protected endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer <access_token>
```

### Token Lifecycle

1. **Access Token**: Short-lived (15 minutes default), used for API requests
2. **Refresh Token**: Long-lived (7 days default), used to get new access tokens

### Token Refresh Flow

```
┌──────────┐          ┌──────────┐          ┌──────────┐
│  Client  │          │   API    │          │   Redis  │
└────┬─────┘          └────┬─────┘          └────┬─────┘
     │                     │                     │
     │ POST /refresh       │                     │
     │ {refresh_token}     │                     │
     │────────────────────>│                     │
     │                     │                     │
     │                     │ Check blacklist     │
     │                     │────────────────────>│
     │                     │                     │
     │                     │<────────────────────│
     │                     │                     │
     │ 200 OK              │                     │
     │ {access_token,      │                     │
     │  refresh_token}     │                     │
     │<────────────────────│                     │
     │                     │                     │
```

---

## Rate Limiting

| Endpoint | Limit |
|----------|-------|
| `/login` | 5 requests per minute |
| `/register` | 3 requests per minute |
| `/forgot-password` | 3 requests per minute |
| All others | 60 requests per minute |

**Rate Limit Headers:**
```
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1706000060
```

**Rate Limit Exceeded Response (429):**
```json
{
  "error": "Too many requests",
  "code": "RATE_LIMIT_EXCEEDED",
  "retry_after": 60
}
```
