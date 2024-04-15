package main

import (
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

	appConfig := storage.LoadAppConfig()

	config := config.LoadConfig()

	ctr := controllers.Controllers{
		Storage: appConfig.Storage,
	}

	router := controllers.InitRoutes(&ctr)
	serverAddress := fmt.Sprintf("%s:%d", config.APIHost, config.APIPort)

	fmt.Printf("Server is running on http://%s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))

	//ctr := controllers.Controllers{
	//	Storage: appConfig.Storage,
	//}
	//
	//router := controllers.InitRoutes(&ctr)
	//
	//fmt.Println("Server is running on http://localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", router))
}
