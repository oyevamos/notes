package routes

import (
	"github.com/oyevamos/notes.git/controllers"
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", controllers.CreateHandler)
	mux.HandleFunc("/read", controllers.ReadHandler)

	return mux
}
