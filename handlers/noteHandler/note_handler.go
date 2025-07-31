package noteHandler

import (
	"lbn/db"
	"lbn/models"
	"lbn/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var noteCollection *mongo.Collection = db.GetCollection(db.DB, "notes")
var validate = validator.New()

type NotesHandler struct {
	DB db.NotesDB
}

func New(db db.NotesDB) *NotesHandler {
	return &NotesHandler{DB: db}
}

func (n *NotesHandler) CreateNote(c echo.Context) error {
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
	insertedNote, err := n.DB.InsertNote(ctx, newNote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Errors: []string{err.Error()}})

	}
	return c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusCreated, Message: "success", Note: insertedNote})
}

func (n *NotesHandler) GetNotes(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	notes, err := n.DB.FindNotes(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NotesResponse{Status: http.StatusInternalServerError, Message: "error", Errors: []string{err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.NotesResponse{Status: http.StatusOK, Message: "success", Notes: notes})
}

func (n *NotesHandler) GetNote(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	note, err := n.DB.FindNote(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusNotFound, Message: "error", Errors: []string{err.Error()}})
	}
	return c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusOK, Message: "success", Note: note})
}

func (n *NotesHandler) DeleteNote(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	note, err := n.DB.DeleteNote(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusNotFound, Message: "error", Errors: []string{err.Error()}})
	}
	return c.JSON(http.StatusCreated, note)
}
