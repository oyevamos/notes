package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/oyevamos/notes.git/config"
	"github.com/oyevamos/notes.git/models"
)

type Controllers struct {
	Config *config.AppConfig
}

func (c *Controllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.InsertNoteIntoDB(r.Context(), &note)
	if err != nil {
		http.Error(w, "Failed to insert note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controllers) InsertNoteIntoDB(ctx context.Context, note *models.Note) error {
	_, err := c.Config.DBPool.Exec(ctx, "INSERT INTO notes (user_id, header, content, date_created, date_modified) VALUES ($1, $2, $3, $4, $5)", note.UserID, note.Header, note.Content, note.DateCreated, note.DateModified)
	if err != nil {
		log.Printf("Error inserting note into database: %v", err)
	}
	return err
}
