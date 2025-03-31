package db

import (
	"context"
	"lbn/models"
)

type NotesDB interface {
	InsertNote(ctx context.Context, note models.Note) (models.Note, error)
	FindNote(ctx context.Context, id string) (models.Note, error)
	FindNotes(ctx context.Context) ([]models.Note, error)
}
