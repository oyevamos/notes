package controllers

import (
	"context"
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

	var content string
	err := c.queryContent(title, &content)
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

func (c *Controllers) queryContent(header string, content *string) error {
	err := appConfig.DBPool.QueryRow(context.Background(), "SELECT content FROM notes WHERE header = $1", header).Scan(content)
	return err
}
