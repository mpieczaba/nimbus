package app

import (
	"log"

	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func New() *App {
	return &App{}
}

func (app *App) Start() {
	// Connect to database
	app.ConnectToDatabase()

	// Start http server
	log.Fatal(app.ServeHTTP())
}
