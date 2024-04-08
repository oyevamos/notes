package main

import (
	"github.com/gorilla/mux"
	"github.com/oyevamos/notes.git/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.NotesRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7500", r))
}
