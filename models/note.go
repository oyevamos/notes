package models

import "time"

type Note struct {
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
}
