package main

import (
	"fmt"
	"log"

	"github.com/Asuzaka/chat-app/backend/pkg/config"
)

func main() {
	cfg := config.Load()

	fmt.Println("Loaded configuration:")
	fmt.Printf("App: %s v%s\n", cfg.App.Name, cfg.App.Version)
	fmt.Printf("Server: %s:%s (%s)\n", "localhost", cfg.Server.Port, cfg.Server.Environment)
	fmt.Printf("DB URL: %s\n", cfg.DatabaseURL())

	log.Println("Backend is running...")
}
