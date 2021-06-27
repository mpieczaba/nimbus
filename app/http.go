package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mpieczaba/nimbus/api/directives"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/filesystem"
	"github.com/mpieczaba/nimbus/store"
	"github.com/mpieczaba/nimbus/validators"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *App) ServeHTTP() error {
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.http = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	app.http.Use(cors.New(corsConfig), auth.Middleware())

	app.http.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nimbus - extensible storage system")
	})

	app.store = store.New(app.db)

	gqlHandler := handler.NewDefaultServer(server.NewExecutableSchema(app.getGQLConfig()))

	app.http.POST("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

	app.http.GET("/playground", func(c *gin.Context) {
		gqlPlayground.ServeHTTP(c.Writer, c.Request)
	})

	app.http.GET("/files/:id", filesystem.SendFile(app.store))

	log.Println("Nimbus server listening on " + os.Getenv("HOST"))

	srv := &http.Server{
		Addr:           os.Getenv("HOST"),
		Handler:        app.http,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv.ListenAndServe()
}

func (app *App) getGQLConfig() server.Config {
	var cfg server.Config

	cfg.Resolvers = &resolvers.Resolver{
		Store:     app.store,
		Validator: validators.New(),
	}

	cfg.Directives.Auth = directives.Auth()
	cfg.Directives.IsAdmin = directives.IsAdmin()

	return cfg
}
