package resolvers

import (
	"github.com/mpieczaba/nimbus/api/server"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/validators"
)

//go:generate go run github.com/99designs/gqlgen

type Store struct {
	User *user.UserStore
	File *file.FileStore
}

type Resolver struct {
	Store     *Store
	Validator *validators.Validator
}

func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
