package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	http *gin.Engine
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
