package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	var userID string

	if id != nil {
		userID = *id
	} else {
		claims, _ := auth.ClaimsFromContext(ctx)

		userID = claims.ID
	}

	return r.Store.User.GetUser("id = ?", userID)
}

func (r *queryResolver) Users(ctx context.Context, after, before *string, first, last *int, username *string) (*models.UserConnection, error) {
	return r.Store.User.GetAllUsers(after, before, first, last, username)
}

// Mutation

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, gqlerror.Errorf("Cannot parse password!")
	}

	return r.Store.User.CreateUser(&models.User{
		ID:       xid.New().String(),
		Username: input.Username,
		Password: string(pass),
		Kind:     models.UserKindUser,
	})
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, input models.UserUpdateInput) (*models.User, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	var userUpdateID string

	if id != nil {
		userUpdateID = *id
	} else {
		claims, _ := auth.ClaimsFromContext(ctx)

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

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*models.User, error) {
	var userDeleteID string

	if id != nil {
		userDeleteID = *id
	} else {
		claims, _ := auth.ClaimsFromContext(ctx)

		userDeleteID = claims.ID
	}

	return r.Store.User.DeleteUser("id = ?", userDeleteID)
}
