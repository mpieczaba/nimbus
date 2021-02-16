package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) Tag(ctx context.Context, id string) (*models.Tag, error) {
	var tag models.Tag

	if err := r.DB.Where("id = ?", id).First(&tag).Error; err != nil {
		return &tag, gqlerror.Errorf("Tag with id `" + id + "` not found!")
	}

	return &tag, nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*models.Tag, error) {
	var tags []*models.Tag

	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all tags!")
	}

	return tags, nil
}

// Field resolver

func (r *tagResolver) Owner(ctx context.Context, obj *models.Tag) (*models.User, error) {
	var owner models.User

	if err := r.DB.Where("id = ?", obj.OwnerID).First(&owner).Error; err != nil {
		return &owner, gqlerror.Errorf("Owner not found")
	}

	return &owner, nil
}
