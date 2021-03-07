package cmd

import (
	"os"

	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/file/file_tag"
	"github.com/mpieczaba/nimbus/filesystem"
	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/tag/tag_share"
	"github.com/mpieczaba/nimbus/user"
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
	fs := filesystem.NewFilesystem()

	// Create data directory if it does not exist
	if err := fs.CreateDataDirectory(); err != nil {
		panic(err)
	}

	// Connect to database
	app.db = database.Connect()

	app.db.AutoMigrate(user.User{}, file.File{}, tag.Tag{}, file_tag.FileTag{}, tag_share.TagShare{}, file_share.FileShare{})

	app.http = fiber.New()

	app.http.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nimbus - extensible storage system with quick access to data")
	})

	// Set up GraphQL api endpoint
	app.http.All("/graphql", func(c *fiber.Ctx) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
			Store: &resolvers.Store{
				User:      user.NewStore(app.db),
				File:      file.NewStore(app.db),
				FileShare: file_share.NewStore(app.db),
				FileTag:   file_tag.NewStore(app.db),
				Tag:       tag.NewStore(app.db),
				TagShare:  tag_share.NewStore(app.db),
			},
			Auth:       auth.NewAuth(c),
			Filesystem: fs,
			Validator:  validators.New(),
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
