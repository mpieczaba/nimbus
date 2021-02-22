package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	var user models.User

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &user, err
	}

	id := claims["id"].(string)

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &user, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &user, gqlerror.Errorf("User with id `" + id + "` not found!")
	}

	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all users!")
	}

	return users, nil
}

// Mutation

func (r *mutationResolver) UserCreate(ctx context.Context, input models.UserInput) (*models.User, error) {
	var user models.User

	if err := r.Validator.Validate(input); err != nil {
		return &user, err
	}

	id := xid.New()

	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return &user, gqlerror.Errorf("Cannot parse password!")
	}

	user = models.User{
		ID:       id.String(),
		Username: input.Username,
		Password: string(pass),
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return &user, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return &user, nil
}

func (r *mutationResolver) UserUpdate(ctx context.Context, input models.UserUpdateInput) (*models.User, error) {
	var user models.User

	if err := r.Validator.Validate(input); err != nil {
		return &user, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &user, err
	}

	id := claims["id"].(string)

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &user, gqlerror.Errorf("User not found!")
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			return &user, gqlerror.Errorf("Cannot parse password!")
		}

		user.Password = string(pass)
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return &user, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return &user, nil
}

func (r *mutationResolver) UserDelete(ctx context.Context) (*models.User, error) {
	var user models.User

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &user, err
	}

	id := claims["id"].(string)

	if err := r.DB.Where("id = ?", id).First(&user).Delete(&user).Error; err != nil {
		return &user, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

// Field resolver

func (r *userResolver) Files(ctx context.Context, obj *models.User) ([]*models.File, error) {
	var files []*models.File

	if err := r.DB.Where("owner_id = ?", obj.ID).Find(&files).Error; err != nil {
		return files, gqlerror.Errorf("Internal database error occurred while getting user files!")
	}

	return files, nil
}
