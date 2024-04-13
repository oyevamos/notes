package controllers

import (
	"encoding/json"
	"github.com/oyevamos/notes.git/models"
	"net/http"
)

func (c *Controllers) ReadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}

	content, err := c.Storage.GetNoteContentByHeader(r.Context(), title)
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
