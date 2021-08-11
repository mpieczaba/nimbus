package handlers

import (
	"github.com/mpieczaba/nimbus/api/directives"
	"github.com/mpieczaba/nimbus/api/resolvers"
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/store"
	"github.com/mpieczaba/nimbus/validators"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func GQLServerHandler(store *store.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler.NewDefaultServer(server.NewExecutableSchema(server.Config{
			Resolvers: &resolvers.Resolver{
				Store:     store,
				Validator: validators.New(),
			},
			Directives: server.DirectiveRoot{
				Auth:    directives.Auth(),
				IsAdmin: directives.IsAdmin(),
			},
		})).ServeHTTP(c.Writer, c.Request)
	}
}

func GQLPlaygroundHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	}
}
