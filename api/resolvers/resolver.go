package resolvers

import (
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/store"
	"github.com/mpieczaba/nimbus/validators"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	Store     *store.Store
	Validator *validators.Validator
}

func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *Resolver) File() server.FileResolver { return &fileResolver{r} }

type fileResolver struct{ *Resolver }

func (r *Resolver) Tag() server.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
