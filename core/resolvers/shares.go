package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// TagShare

// Field resolver

func (r *tagShareResolver) User(ctx context.Context, obj *models.TagShare) (*models.User, error) {
	var user models.User

	if err := r.DB.Where("id = ?", obj.UserID).First(&user).Error; err != nil {
		return &user, gqlerror.Errorf("Internal database error occurred while getting user!")
	}

	return &user, nil
}
