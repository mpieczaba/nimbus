package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/api/models"
	"github.com/mpieczaba/nimbus/user"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	return r.Store.User.GetAllUsers()
}

// Mutation

func (r *mutationResolver) CreateUser(ctx context.Context, input user.UserInput) (*user.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, gqlerror.Errorf("Cannot parse password!")
	}

	return r.Store.User.CreateUser(&user.User{
		ID:       xid.New().String(),
		Username: input.Username,
		Password: string(pass),
		Kind:     models.UserKindUser,
	})
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, input user.UserUpdateInput) (*user.User, error) {
	return nil, gqlerror.Errorf("Not implemented!")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*user.User, error) {
	return nil, gqlerror.Errorf("Not implemented!")
}
