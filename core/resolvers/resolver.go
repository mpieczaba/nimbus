package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/mpieczaba/nimbus/core/generated"
	"github.com/mpieczaba/nimbus/core/validators"

	"gorm.io/gorm"
)

type Resolver struct {
	DB        *gorm.DB
	Validator *validators.Validator
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
