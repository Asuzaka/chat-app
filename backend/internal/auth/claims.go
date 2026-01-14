package auth

import "github.com/golang-jwt/jwt/v5"

type AccessClaims struct {
	UserID   string `json:"sub"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}
