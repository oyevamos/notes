package storage

import (
	"context"
	"fmt"
	"github.com/oyevamos/notes.git/config"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppConfig struct {
	Storage *Storage
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
		log.Fatalf("Unable to connect to database: %v", err)
	}
	storage := NewStorage(dbpool)
	return &AppConfig{
		Storage: storage,
	}, nil
}
