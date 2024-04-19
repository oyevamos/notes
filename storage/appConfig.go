package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oyevamos/notes.git/config"
)

type AppConfig struct {
	Storage *Storage
}

func (s *Storage) Ping(ctx context.Context) error {
	return s.DBPool.Ping(ctx)
}

func LoadAppConfig() (*AppConfig, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	dbURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	storage := NewStorage(dbpool)
	return &AppConfig{
		Storage: storage,
	}, nil
}
