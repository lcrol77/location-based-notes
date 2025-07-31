package mongoImpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lbn/models"
	"log"
)

type NotesMongoDatabase struct {
	Collection *mongo.Collection
}

func (m *NotesMongoDatabase) InsertNote(ctx context.Context, newNote models.Note) (models.Note, error) {
	var note models.Note
	result, err := m.Collection.InsertOne(ctx, newNote)
	if err != nil {
		return note, err
	}
	err = m.Collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&note)
	return note, err
}

func (m *NotesMongoDatabase) FindNotes(ctx context.Context) ([]models.Note, error) {
	var notes []models.Note
	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return notes, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &notes); err != nil {
		return notes, err
	}
	return notes, err
}

func (m *NotesMongoDatabase) FindNote(ctx context.Context, id string) (models.Note, error) {
	var note models.Note
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("Invalid ObjectID: ", err)
		return note, err
	}
	err = m.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&note)
	if err != nil {
		log.Print("Note not found: ", err)
		return note, err
	}
	return note, err
}

func (m *NotesMongoDatabase) DeleteNote(ctx context.Context, id string) (mongo.DeleteResult, error) {
	var res mongo.DeleteResult
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("Invalid ObjectID: ", err)
		return res, err
	}
	ptr, err := m.Collection.DeleteOne(ctx, bson.M{"_id": objId})	
	if err != nil {
		log.Print("Note not found: ", err)
		return res, err
	}
	res = *ptr
	return res, err
}
