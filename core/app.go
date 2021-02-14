package core

import (
	"os"

	"github.com/mpieczaba/nimbus/core/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	db   *gorm.DB
	http *fiber.App
}

func NewApp() *App {
	app := &App{}

	return app
}

func (app *App) Start() {
	// Connect to database
	app.db = database.Connect()

	app.http = fiber.New()

	app.http.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nimbus - extensible storage system with quick access to data")
	})

	if err := app.http.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
