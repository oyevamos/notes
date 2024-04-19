package controllers

import (
	"encoding/json"
	"github.com/oyevamos/notes.git/models"
	"net/http"
)

func (c *Controllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.Storage.InsertNote(r.Context(), &note)
	if err != nil {
		http.Error(w, "Failed to insert note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
