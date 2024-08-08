package main

import (
	"log"
	"notes/db"
	"notes/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/allNotes", handlers.GetAllNotes)
	e.GET("/notes", handlers.GetNotes) // get notes within a specific area. has lat and long in params
	e.POST("/note", handlers.CreateNote)
	e.PUT("/note/:id", handlers.UpdateNote)
	e.DELETE("/note/:id", handlers.DeleteNote)
	db.InitMongodbClient()
	defer db.CloseClientConnection()
	e.Logger.Fatal(e.Start(":8080"))
}
