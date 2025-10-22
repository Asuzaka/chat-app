package db

import (
	"context"
	"time"

	"github.com/Asuzaka/chat-app/backend/pkg/config"
	"github.com/Asuzaka/chat-app/backend/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init(config config.Config) error {
	cfg, err := pgxpool.ParseConfig(config.DatabaseURL())
	if err != nil {
		return err
	}

	cfg.MaxConns = int32(config.Database.MaxOpenConns)
	cfg.MinConns = int32(config.Database.MaxIdleConns)
	cfg.MaxConnLifetime = config.Database.ConnMaxLifetime
	cfg.HealthCheckPeriod = 1 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	Pool = pool

	logger.Info("PostgreSQL connection pool initialized successfully")

	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
		logger.Info("Database pool closed gracefully")
	}
}
