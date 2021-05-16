package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/api/models"
	"github.com/mpieczaba/nimbus/auth"
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
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

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
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.GetAuthClaimsFromContext(ctx)

	var userUpdateID string

	if id != nil {
		userUpdateID = *id
	} else {
		userUpdateID = claims.ID
	}

	userToUpdate, err := r.Store.User.GetUser("id = ?", userUpdateID)

	if err != nil {
		return nil, err
	}

	if input.Username != "" {
		userToUpdate.Username = input.Username
	}

	if input.Password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, gqlerror.Errorf("Cannot parse password!")
		}

		userToUpdate.Password = string(pass)
	}

	if input.Kind != "" {
		userToUpdate.Kind = input.Kind
	}

	return r.Store.User.UpdateUser(userToUpdate)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*user.User, error) {
	claims, _ := auth.GetAuthClaimsFromContext(ctx)

	var userDeleteID string

	if id != nil {
		userDeleteID = *id
	} else {
		userDeleteID = claims.ID
	}

	return r.Store.User.DeleteUser("id = ?", userDeleteID)
}
