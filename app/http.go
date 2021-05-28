package app

import (
	"log"
	"net/http"
	"os"

	"github.com/mpieczaba/nimbus/api/directives"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/store"
	"github.com/mpieczaba/nimbus/validators"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (app *App) ServeHTTP() error {
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.http = gin.Default()

	app.http.Use(auth.Middleware())

	app.http.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nimbus - extensible storage system")
	})

	var cfg server.Config

	cfg.Resolvers = &resolvers.Resolver{
		Store:     store.New(app.db),
		Validator: validators.New(),
	}

	cfg.Directives.Auth = directives.Auth()
	cfg.Directives.IsAdmin = directives.IsAdmin()

	gqlHandler := handler.NewDefaultServer(server.NewExecutableSchema(cfg))

	gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

	app.http.POST("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	app.http.GET("/playground", func(c *gin.Context) {
		gqlPlayground.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Nimbus server listening on http://127.0.0.1:" + os.Getenv("PORT"))

	return app.http.Run(":" + os.Getenv("PORT"))
}
