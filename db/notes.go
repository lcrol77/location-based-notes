package db

import (
	"context"
	"notes/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllNotes() ([]models.Note, error) {
	coll := GetDB().Collection("notes")
	notes := []models.Note{}
    findOptions := options.Find()
    cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return notes, err
	}
	cursor.All(context.TODO(), &notes)
	return notes, nil
}
