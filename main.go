package main

import (
	"fmt"
	"github.com/oyevamos/notes.git/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/oyevamos/notes.git/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig := storage.LoadAppConfig()
	ctr := controllers.Controllers{Config: appConfig}

	router := controllers.InitRoutes(&ctr)

	controllers.SetAppConfig(appConfig)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
