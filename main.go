package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/oyevamos/notes.git/config"
	"github.com/oyevamos/notes.git/controllers"
	"github.com/oyevamos/notes.git/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig := config.LoadAppConfig()
	ctr := controllers.Controllers{Config: appConfig}

	router := routes.InitRoutes(&ctr)

	controllers.SetAppConfig(appConfig)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
