package handlers

import (
	"net/http"
	"notes/db"
	"notes/models"

	"github.com/labstack/echo/v4"
)

func GetAllNotes(c echo.Context) error {
	return nil
}
func GetNotes(c echo.Context) error {
	return nil
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
	return nil
}
func DeleteNote(c echo.Context) error {
	return nil
}
