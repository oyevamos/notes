// config.go

package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppConfig struct {
	DBPool *pgxpool.Pool
}

func LoadAppConfig() *AppConfig {
	dbPool := ConnectDB()
	return &AppConfig{
		DBPool: dbPool,
	}
}

func ConnectDB() *pgxpool.Pool {
	cfg := LoadConfig()
	dbURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return dbpool
}
