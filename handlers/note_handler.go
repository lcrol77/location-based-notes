package handlers

import (
	"lbn/db"
	"lbn/models"
	"lbn/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&note); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	newNote := models.Note{
		Id:    primitive.NewObjectID(),
		Title: note.Title,
		Body:  note.Body,
	}
	result, err := noteCollection.InsertOne(ctx, newNote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})

}

func GetNotes(c echo.Context) error {
	return nil
}
