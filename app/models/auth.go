package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// Struct User sesuai tabel users
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Request body untuk login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response login (user info + token)
type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// Payload JWT
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
