package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dbURL := "postgres://buba:12345@localhost:5432/notes_db"
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return dbpool
}
