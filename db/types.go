package db

import (
	"context"
	"lbn/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type NotesDB interface {
	InsertNote(ctx context.Context, note models.Note) (models.Note, error)
	FindNote(ctx context.Context, id string) (models.Note, error)
	FindNotes(ctx context.Context) ([]models.Note, error)
	DeleteNote(ctx context.Context, id string) (mongo.DeleteResult, error)
}
