package routes

import (
	"github.com/gorilla/mux"
	"github.com/oyevamos/notes.git/controllers"
)

var NotesRoutes = func(router *mux.Router) {
	router.HandleFunc("/note", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/note", controllers.GetAllNotes).Methods("GET")
}
