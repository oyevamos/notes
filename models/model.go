package models

import (
	"github.com/jinzhu/gorm"
	"github.com/oyevamos/notes.git/config"
)

var db *gorm.DB

type Note struct {
	Id            int    `json:"id"`
	User_id       int    `json:"user_id"`
	Header        string `json:"header"`
	Content       string `json:"content"`
	Date_created  string `json:"date_created"`
	Date_modified string `json:"date_modified"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Note{})
}

func (n *Note) CreateNote() *Note {
	db.NewRecord(n)
	db.Create(&n)
	return n
}

func GetAllNotes() []Note {
	var Notes []Note
	db.Find(&Notes)
	return Notes
}
