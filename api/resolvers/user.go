package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/user"
)

// Query

func (r *queryResolver) User(ctx context.Context) (*user.User, error) {
	return nil, nil
}

// Mutation

func (r *mutationResolver) CreateUser(ctx context.Context) (*user.User, error) {
	return nil, nil
}
