package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(ctr *Controllers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create", ctr.CreateHandler).Methods(http.MethodPost)
	router.HandleFunc("/read", ctr.ReadHandler).Methods(http.MethodGet)

	return router
}
