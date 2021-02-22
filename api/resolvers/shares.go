package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/user"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// TagShare

// Field resolver

func (r *tagShareResolver) User(ctx context.Context, obj *models.TagShare) (*user.User, error) {
	var tagUser user.User

	if err := r.DB.Where("id = ?", obj.UserID).First(&tagUser).Error; err != nil {
		return &tagUser, gqlerror.Errorf("Internal database error occurred while getting user!")
	}

	return &tagUser, nil
}

func (r *fileShareResolver) User(ctx context.Context, obj *models.FileShare) (*user.User, error) {
	var fileUser user.User

	if err := r.DB.Where("id = ?", obj.UserID).First(&fileUser).Error; err != nil {
		return &fileUser, gqlerror.Errorf("Internal database error occurred while getting user!")
	}

	return &fileUser, nil
}
