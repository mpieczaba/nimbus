package app

import (
	"log"

	"github.com/mpieczaba/nimbus/user"

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

	app.db.AutoMigrate(user.User{})

	// Start http server
	log.Fatal(ServeHTTP())
}
