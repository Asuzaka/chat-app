package main

import (
	"fmt"

	"github.com/Asuzaka/chat-app/backend/pkg/config"
	"github.com/Asuzaka/chat-app/backend/pkg/db"
	"github.com/Asuzaka/chat-app/backend/pkg/logger"
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

	if err := db.Init(*cfg); err != nil {
		logger.Error("Failed to initialize database pool: " + err.Error())
		return
	}

	defer db.Close()

	logger.Info("Backend is running...")
}
