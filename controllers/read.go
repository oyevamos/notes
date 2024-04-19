package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (c *Controllers) ReadHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID Format", http.StatusBadRequest)
		return
	}

	note, err := c.Storage.GetNoteById(r.Context(), id)
	if err != nil {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
