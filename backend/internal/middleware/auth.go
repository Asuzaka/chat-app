package middleware

import (
	"github.com/Asuzaka/chat-app/backend/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func Auth(jwtManager *auth.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		tokenStr := authHeader[len("Bearer "):]

		if tokenStr == authHeader {
			return fiber.ErrUnauthorized
		}

		claims, err := jwtManager.VerifyAccessToken(tokenStr)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// Store user information in context locals
		c.Locals("userID", claims.UserID)

		return c.Next()
	}
}
