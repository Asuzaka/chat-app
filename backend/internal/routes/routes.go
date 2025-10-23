package routes

import (
	"github.com/Asuzaka/chat-app/backend/internal/health"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/health", health.Handler)

	// other routes and ws
}
