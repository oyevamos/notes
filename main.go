package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oyevamos/notes.git/config"
	"github.com/oyevamos/notes.git/controllers"
	"github.com/oyevamos/notes.git/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool := config.ConnectDB()
	defer dbpool.Close()

	// Установка dbpool для контроллеров
	controllers.SetDBPool(dbpool)

	router := routes.InitRoutes()

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
