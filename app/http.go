package app

import (
	"log"
	"os"

	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (app *App) ServeHTTP() error {
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	http := gin.Default()

	http.Use(auth.Middleware())

	http.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("Nimbus - extensible storage system")
	})

	var cfg generated.Config

	cfg.Resolvers = &resolvers.Resolver{
		Store: &resolvers.Store{
			User: user.NewUserStore(app.db),
		},
	}

	cfg.Directives.Auth = auth.Directive()

	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

	http.POST("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	http.GET("/playground", func(c *gin.Context) {
		gqlPlayground.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Nimbus server listening on http://127.0.0.1:" + os.Getenv("PORT"))

	return http.Run(":" + os.Getenv("PORT"))
}
