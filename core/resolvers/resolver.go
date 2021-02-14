package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/mpieczaba/nimbus/core/generated"

	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
