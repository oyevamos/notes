package config

import (
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

func LoadConfig() (Config, error) {
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

	dbPort, err := getIntValue("DB_PORT", 5432)
	if err != nil {
		//log.Fatalf("Error parsing DB_PORT: %s", err)
		return Config{}, err
	}
	apiPort, err := getIntValue("API_PORT", 8080)
	if err != nil {
		//log.Fatalf("Error parsing API_PORT: %s", err)
		return Config{}, err
	}

	return Config{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		APIPort:    apiPort,
		APIHost:    apiHost,
	}, nil
}

func getIntValue(envKey string, defaultValue int) (int, error) {
	if value, exists := os.LookupEnv(envKey); exists {
		port, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		return port, nil
	}
	return defaultValue, nil
}
