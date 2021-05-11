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
	app.db = ConnectToDatabase()

	app.db.AutoMigrate()

	// Start http server
	log.Fatal(ServeHTTP())
}
