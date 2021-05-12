package resolvers

import "github.com/mpieczaba/nimbus/api/generated"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
