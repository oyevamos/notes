package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oyevamos/notes.git/models"
	"log"
)

type Storage struct {
	DBPool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		DBPool: pool,
	}
}

func (s *Storage) InsertNote(ctx context.Context, note *models.Note) error {
	_, err := s.DBPool.Exec(ctx, "INSERT INTO notes (user_id, title, content, date_created, date_modified) VALUES ($1, $2, $3, $4, $5)",
		note.UserID, note.Title, note.Content, note.DateCreated, note.DateModified)
	if err != nil {
		log.Printf("Error inserting note into database: %v", err)
	}
	return err
}

func (s *Storage) GetNoteContentByTitle(ctx context.Context, title string) (string, error) {
	var content string
	err := s.DBPool.QueryRow(ctx, "SELECT content FROM notes WHERE title = $1", title).Scan(&content)
	if err != nil {
		return "", err
	}
	return content, nil
}
