package routes

import (
	"github.com/gorilla/mux"
	"github.com/oyevamos/notes.git/controllers"
)

func InitRoutes(ctr *controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	// Set up routes with handlers that are methods of Controllers
	router.HandleFunc("/create", ctr.CreateHandler).Methods("POST")
	router.HandleFunc("/read", ctr.ReadHandler).Methods("GET")

	return router
}
