package models

import "time"

type Note struct {
	UserID       int       `json:"user_id"`
	Header       string    `json:"header"`
	Content      string    `json:"content"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
}
