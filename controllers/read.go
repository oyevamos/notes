package controllers

import (
	"encoding/json"
	"github.com/oyevamos/notes.git/models"
	"net/http"
)

func (c *Controllers) ReadHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	content, err := c.Storage.GetNoteContentByTitle(r.Context(), id)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	response := models.NoteContent{Content: content}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
