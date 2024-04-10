package main

import (
	"fmt"
	"github.com/oyevamos/notes.git/config"
	"github.com/oyevamos/notes.git/controllers"
	"github.com/oyevamos/notes.git/routes"
	"log"
	"net/http"
)

func main() {
	dbpool := config.ConnectDB()
	defer dbpool.Close()

	// Установка dbpool для контроллеров
	controllers.SetDBPool(dbpool)

	router := routes.InitRoutes()

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
