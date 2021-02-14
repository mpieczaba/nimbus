package routes

import (
	"github.com/mpieczaba/nimbus/core/generated"
	"github.com/mpieczaba/nimbus/core/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GraphQL(router fiber.Router, db *gorm.DB) {
	router.All("/graphql", func(c *fiber.Ctx) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
			DB: db,
		}}))

		gqlHandler := srv.Handler()

		gqlHandler(c.Context())

		return nil
	})

	router.Get("/playground", func(c *fiber.Ctx) error {
		gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

		gqlPlayground(c.Context())

		return nil
	})
}
