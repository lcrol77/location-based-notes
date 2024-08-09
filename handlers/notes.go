package handlers

import (
	"net/http"
	"notes/db"
	"notes/models"

	"github.com/labstack/echo/v4"
)

func GetAllNotes(c echo.Context) error {
	notes, err := db.GetAllNotes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, notes)
}

func GetNotes(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented yet")
}

func GetNote(c echo.Context) error {
	var id string
	c.Bind(&id)
	note, err:= db.GetNote(id);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, note)
}

func CreateNote(c echo.Context) error {
	note := models.Note{}
	c.Bind(&note)
	newNote, err := db.CreateNote(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newNote)
}

func UpdateNote(c echo.Context) error {
	note := models.Note{}
	c.Bind(&note)
	newNote, err := db.UpdateNote(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newNote)
}

func DeleteNote(c echo.Context) error {
	note := models.Note{}
	c.Bind(&note)
	newNote, err := db.DeleteNote(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newNote)
}
