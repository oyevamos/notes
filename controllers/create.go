package controllers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oyevamos/notes.git/models"
	"net/http"
)

var dbpool *pgxpool.Pool

func SetDBPool(pool *pgxpool.Pool) {
	dbpool = pool
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := dbpool.Exec(context.Background(), "INSERT INTO notes (user_id, header, content, date_created, date_modified) VALUES ($1, $2, $3, $4, $5)", note.UserID, note.Header, note.Content, note.DateCreated, note.DateModified)
	if err != nil {
		http.Error(w, "Failed to insert note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
