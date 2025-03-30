package handlers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lbn/db"
	"lbn/models"
	"lbn/responses"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var noteCollection *mongo.Collection = db.GetCollection(db.DB, "notes")
var validate = validator.New()

func CreateNote(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var note models.Note
	defer cancel()

	//validate the request body
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "errors", Errors: []string{err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&note); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Errors: []string{validationErr.Error()}})
	}
	newNote := models.Note{
		Title: note.Title,
		Body:  note.Body,
	}
	result, err := noteCollection.InsertOne(ctx, newNote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Errors: []string{err.Error()}})
	}
	err = noteCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&note)
	return c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusCreated, Message: "success", Note: note})
}

func GetNotes(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := noteCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NotesResponse{Status: http.StatusInternalServerError, Message: "error", Errors: []string{err.Error()}})
	}
	defer cursor.Close(ctx)

	// Store results
	var notes []models.Note
	if err = cursor.All(ctx, &notes); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NotesResponse{Status: http.StatusInternalServerError, Message: "error", Errors: []string{err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.NotesResponse{Status: http.StatusOK, Message: "success", Notes: notes})
}

func GetNote(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("Invalid ObjectID: ", err)
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Errors: []string{err.Error()}})
	}
	var note models.Note
	err = noteCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&note)
	if err != nil {
		log.Print("Note not found: ", err)
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusNotFound, Message: "error", Errors: []string{err.Error()}})
	}
	return c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusOK, Message: "success", Note: note})
}
