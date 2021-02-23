package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

func (r *queryResolver) Me(ctx context.Context) (*user.User, error) {
	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	return r.UserStore.GetUserById(claims["id"].(string))
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserStore.GetUserById(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	return r.UserStore.GetAllUsers()
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

	return r.UserStore.SaveUser(&user.User{
		ID:       xid.New().String(),
		Username: input.Username,
		Password: string(pass),
	})
}

func (r *mutationResolver) UserUpdate(ctx context.Context, input user.UserUpdateInput) (*user.User, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	userToUpdate, err := r.UserStore.GetUserById(claims["id"].(string))

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

	return r.UserStore.SaveUser(userToUpdate)
}

func (r *mutationResolver) UserDelete(ctx context.Context) (*user.User, error) {
	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	return r.UserStore.DeleteUser(claims["id"].(string))
}

// Field resolver

func (r *userResolver) Files(ctx context.Context, obj *user.User) ([]*file.File, error) {
	var files []*file.File

	if err := r.DB.Where("owner_id = ?", obj.ID).Find(&files).Error; err != nil {
		return files, gqlerror.Errorf("Internal database error occurred while getting user files!")
	}

	return files, nil
}
