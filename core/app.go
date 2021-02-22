package core

import (
	"os"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/core/routes"
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/utils"

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
	// Create data directory if it does not exist
	if err := utils.CreateDataDirectory(); err != nil {
		panic(err)
	}

	// Connect to database
	app.db = database.Connect()

	app.db.AutoMigrate(models.User{}, models.File{}, models.Tag{}, models.FileTag{}, models.TagShare{}, models.FileShare{})

	app.http = fiber.New()

	app.http.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nimbus - extensible storage system with quick access to data")
	})

	// GraphQL api endpoint and playground
	routes.GraphQL(app.http, app.db)

	if err := app.http.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
