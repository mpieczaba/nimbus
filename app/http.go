package app

import (
	"log"
	"os"

	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/api/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
)

func ServeHTTP() error {
	http := fiber.New(fiber.Config{DisableStartupMessage: true})

	http.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nimbus - extensible storage system")
	})

	http.All("/graphql", func(c *fiber.Ctx) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

		gqlHandler := srv.Handler()

		gqlHandler(c.Context())

		return nil
	})

	http.Get("/playground", func(c *fiber.Ctx) error {
		gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

		gqlPlayground(c.Context())

		return nil
	})

	log.Println("Nimbus server listening on http://127.0.0.1:" + os.Getenv("PORT"))

	return http.Listen(":" + os.Getenv("PORT"))
}
