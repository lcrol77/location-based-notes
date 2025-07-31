package routes

import (
	"github.com/labstack/echo/v4"
	"lbn/db"
	"lbn/db/mongoImpl"
	"lbn/handlers/noteHandler"
)

func NotesRoute(e *echo.Echo) {
	n := noteHandler.New(&mongoImpl.NotesMongoDatabase{Collection: db.GetCollection(db.DB, "notes")})
	e.GET("/notes", n.GetNotes)
	e.GET("/note/:id", n.GetNote)
	e.DELETE("/note/:id", n.DeleteNote)
	e.POST("/note", n.CreateNote)
}
