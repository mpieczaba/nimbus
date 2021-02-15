package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/mpieczaba/nimbus/core/generated"
	"github.com/mpieczaba/nimbus/core/validators"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Resolver struct {
	Ctx       *fiber.Ctx
	DB        *gorm.DB
	Validator *validators.Validator
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
