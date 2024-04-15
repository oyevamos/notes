package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	APIHost    string
	APIPort    int
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() Config {

	getPort := func(envKey string, defaultValue int) int {
		if value, exists := os.LookupEnv(envKey); exists {
			port, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("Error parsing %s: %s", envKey, err)
			}
			return port
		}
		return defaultValue
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "buba"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "12345"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "notes_db"
	}
	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		apiHost = "localhost"
	}

	return Config{
		DBHost:     dbHost,
		DBPort:     getPort("DB_PORT", 5432),
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		APIPort:    getPort("API_PORT", 8080),
		APIHost:    apiHost,
	}
}
