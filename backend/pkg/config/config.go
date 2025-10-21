package config

import (
	"fmt"
	"os"
	"time"

	"github.com/Asuzaka/chat-app/backend/pkg/logger"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name     string `yaml:"name"`
		Version  string `yaml:"version"`
		Timezone string `yaml:"timezone"`
	} `yaml:"app"`

	Server struct {
		Port        string        // .env
		ReadTimeout time.Duration `yaml:"read_timout"`
		WriteTimout time.Duration `yaml:"write_timout"`
		Environment string        //.env
	} `yaml:"server"`

	Database struct {
		Host            string        //.env
		Port            string        //.env
		User            string        //.env
		Password        string        //.env
		Name            string        //.env
		SSLMode         string        //.env
		MaxOpenConns    int           `yaml:"max_open_conns"`
		MaxIdleConns    int           `yaml:"max_idle_conns"`
		ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	} `yaml:"database"`

	JWT struct {
		Secret string //.env
	}
}

func Load() *Config {
	// Load .env file variables
	_ = godotenv.Load()

	cfg := &Config{}

	// Load config.yaml file variables
	yamlFile, err := os.ReadFile("config/config.yaml")

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read config.yaml: %v", err))
	}

	if err = yaml.Unmarshal(yamlFile, cfg); err != nil {
		logger.Error(fmt.Sprintf("Failed to parse config.yaml: %v", err))
	}

	// Insert values to config
	cfg.Server.Port = getEnv("PORT", "8080")
	cfg.Server.Environment = getEnv("ENV", "development")

	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")
	cfg.Database.Name = getEnv("DB_NAME", "chatapp")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	cfg.JWT.Secret = getEnv("JWT_SECRET", "unknownsecret")
	return cfg
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
