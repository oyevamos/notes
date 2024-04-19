package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oyevamos/notes.git/config"
	"github.com/oyevamos/notes.git/controllers"
	"github.com/oyevamos/notes.git/storage"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig, nil := storage.LoadAppConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	if err := appConfig.Storage.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	ctr := controllers.Controllers{
		Storage: appConfig.Storage,
	}

	router := controllers.InitRoutes(&ctr)
	serverAddress := fmt.Sprintf("%s:%d", cfg.APIHost, cfg.APIPort)

	fmt.Printf("Server is running on http://%s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))

}
