package cmd

import (
	"os"

	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"
	"github.com/mpieczaba/nimbus/validators"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	app.db.AutoMigrate(user.User{}, file.File{}, tag.Tag{}, file.FileTag{}, tag.TagShare{}, file.FileShare{})

	app.http = fiber.New()

	app.http.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nimbus - extensible storage system with quick access to data")
	})

	// Set up GraphQL api endpoint
	app.http.All("/graphql", func(c *fiber.Ctx) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
			Ctx:       c,
			UserStore: user.NewStore(app.db),
			FileStore: file.NewStore(app.db),
			TagStore:  tag.NewStore(app.db),
			Validator: validators.New(),
		}}))

		gqlHandler := srv.Handler()

		gqlHandler(c.Context())

		return nil
	})

	// Set up GraphQL playground
	app.http.Get("/playground", func(c *fiber.Ctx) error {
		gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

		gqlPlayground(c.Context())

		return nil
	})

	if err := app.http.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
