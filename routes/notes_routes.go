package routes

import (
	"github.com/labstack/echo/v4"
	"lbn/handlers"
)

func NotesRoute(e *echo.Echo) {
	e.GET("/notes", handlers.GetNotes)
	e.POST("/note", handlers.CreateNote)
}
