package db

import (
	"context"
	"notes/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(note models.Note) (models.Note, error) {
	coll := GetDB().Collection("notes")
	result, err := coll.InsertOne(context.TODO(), note)
	note.Id = result.InsertedID.(primitive.ObjectID)
	if err != nil {
		return note, err
	}
	return note, nil
}
