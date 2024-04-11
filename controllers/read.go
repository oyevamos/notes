package controllers

import (
	"context"
	"encoding/json"
	"net/http"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	header := r.URL.Query().Get("header")
	if header == "" {
		http.Error(w, "Missing header parameter", http.StatusBadRequest)
		return
	}

	var content string
	err := dbpool.QueryRow(context.Background(), "SELECT content FROM notes WHERE header = $1", header).Scan(&content)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Content string `json:"content"`
	}{Content: content})
}