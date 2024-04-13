package routes

import (
	"github.com/gorilla/mux"
	"github.com/oyevamos/notes.git/controllers"
	"net/http"
)

func InitRoutes(ctr *controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create", ctr.CreateHandler).Methods(http.MethodPost)
	router.HandleFunc("/read", ctr.ReadHandler).Methods(http.MethodGet)

	return router
}
