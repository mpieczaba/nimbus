package resolvers

import (
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/user"
)

//go:generate go run github.com/99designs/gqlgen

type Store struct {
	User *user.UserStore
}

type Resolver struct {
	Store *Store
}

func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
