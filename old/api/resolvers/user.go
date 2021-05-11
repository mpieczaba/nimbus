package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

func (r *queryResolver) Me(ctx context.Context) (*user.User, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	return r.Store.User.GetUser("id = ?", claims["id"].(string))
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	return r.Store.User.GetAllUsers()
}

// Mutation

func (r *mutationResolver) UserCreate(ctx context.Context, input user.UserInput) (*user.User, error) {
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
		Kind:     "User",
	})
}

func (r *mutationResolver) UserUpdate(ctx context.Context, id *string, input user.UserUpdateInput) (*user.User, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	// Check if user is banned
	if claims["kind"].(string) == "Banned" {
		return nil, gqlerror.Errorf("You have no permissions to update user!")
	}

	var userUpdateID string

	if id != nil && claims["kind"].(string) == "Admin" {
		userUpdateID = *id
	} else {
		userUpdateID = claims["id"].(string)
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

	if input.Kind != "" && claims["kind"].(string) == "Admin" {
		userToUpdate.Kind = input.Kind
	}

	return r.Store.User.UpdateUser(userToUpdate)
}

func (r *mutationResolver) UserDelete(ctx context.Context, id *string) (*user.User, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	// Check if user is banned
	if claims["kind"].(string) == "Banned" {
		return nil, gqlerror.Errorf("You have no permissions to delete user!")
	}

	var userUpdateID string

	if id != nil && claims["kind"].(string) == "Admin" {
		userUpdateID = *id
	} else {
		userUpdateID = claims["id"].(string)
	}

	return r.Store.User.DeleteUser("id = ?", userUpdateID)
}

// Field resolver

func (r *userResolver) Files(ctx context.Context, obj *user.User) ([]*file.File, error) {
	return r.Store.File.GetAllFilesWithCondition("owner_id = ?", obj.ID)
}
