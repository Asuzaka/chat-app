package main

import (
	"fmt"

	"github.com/Asuzaka/chat-app/backend/internal/auth"
	"github.com/Asuzaka/chat-app/backend/internal/middleware"
	"github.com/Asuzaka/chat-app/backend/internal/routes"
	"github.com/Asuzaka/chat-app/backend/pkg/config"
	"github.com/Asuzaka/chat-app/backend/pkg/db"
	"github.com/Asuzaka/chat-app/backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Temporar logger
	logger.Init(true)
	logger.Info("Temporar logger initialized")

	cfg := config.Load()

	logger.Init(cfg.Server.Environment == "development")

	logger.Info("Loaded configuration:")
	logger.Info(fmt.Sprintf("App: %s v%s\n", cfg.App.Name, cfg.App.Version))
	logger.Info(fmt.Sprintf("Server: %s:%s (%s)\n", "localhost", cfg.Server.Port, cfg.Server.Environment))
	logger.Info(fmt.Sprintf("DB URL: %s\n", cfg.DatabaseURL()))

	// Initialize JWT Manager

	jwtmanager := auth.JWTManager{
		AccessSecret:  []byte(cfg.JWT.AccessSecret),
		RefreshSecret: []byte(cfg.JWT.RefreshSecret),
		AcessTTL:      cfg.JWT.AccessTTL,
		RefreshTTL:    cfg.JWT.RefreshTTL,
		Issuer:        cfg.JWT.Issuer,
	}

	logger.Info("JWT Manager initialized.")

	if err := db.Init(*cfg); err != nil {
		logger.Error("Failed to initialize database pool: " + err.Error())
		return
	}

	defer db.Close()

	app := fiber.New()
	app.Use(middleware.Auth(&jwtmanager))
	// REST + WS
	routes.Register(app)

	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		logger.Error(fmt.Sprint(err))
	}

	logger.Info("Backend is running...")
}
