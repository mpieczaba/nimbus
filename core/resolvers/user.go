package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
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
