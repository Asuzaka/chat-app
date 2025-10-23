package health

import (
	"context"
	"time"

	"github.com/Asuzaka/chat-app/backend/pkg/db"
	"github.com/gofiber/fiber/v2"
)

var startTime = time.Now()

func Handler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dbStatus := "connected"
	if db.Pool == nil {
		dbStatus = "not_initialized"
	} else if err := db.Pool.Ping(ctx); err != nil {
		dbStatus = "disconnected"
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"database": dbStatus,
		"uptime":   time.Since(startTime).Truncate(time.Second).String(),
	})
}
