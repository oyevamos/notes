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
	// Преобразование порта базы данных из строки в целое число
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %s", err)
	}

	// Преобразование порта API из строки в целое число
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error parsing API_PORT: %s", err)
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		APIPort:    apiPort,
		APIHost:    os.Getenv("API_HOST"),
	}
}
