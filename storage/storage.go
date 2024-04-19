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

func (s *Storage) GetNoteById(ctx context.Context, id int) (*models.Note, error) {
	var note models.Note
	err := s.DBPool.QueryRow(ctx, "SELECT user_id, title, content, date_created, date_modified FROM notes WHERE id = $1", id).Scan(&note.UserID, &note.Title, &note.Content, &note.DateCreated, &note.DateModified)
	if err != nil {
		return nil, err
	}
	return &note, nil
}
