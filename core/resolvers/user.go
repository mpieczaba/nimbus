package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Query

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

// TODO: Add user updating and deleting (auth required!)
