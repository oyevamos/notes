package controllers

import (
	"encoding/json"
	"github.com/oyevamos/notes.git/models"
	"github.com/oyevamos/notes.git/utils"
	"net/http"
)

var NewNote models.Note

func CreateNote(w http.ResponseWriter, r *http.Request) {
	CreateNote := &models.Note{}
	utils.ParseBody(r, CreateNote)
	n := CreateNote.CreateNote()
	res, _ := json.Marshal(n)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	newNotes := models.GetAllNotes()
	res, _ := json.Marshal(newNotes)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
