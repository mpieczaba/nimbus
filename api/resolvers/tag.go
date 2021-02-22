package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
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

// Mutation

func (r *mutationResolver) TagCreate(ctx context.Context, input models.TagInput) (*models.Tag, error) {
	var tag models.Tag

	if err := r.Validator.Validate(input); err != nil {
		return &tag, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tag, err
	}

	ownerID := claims["id"].(string)

	id := xid.New()

	if len(input.SharedFor) > 0 {
		// Save tag shares
		tagShares := utils.TagShareInputsToTagShares(id.String(), input.SharedFor)

		if err := r.DB.Save(&tagShares).Error; err != nil {
			return &tag, gqlerror.Errorf("Cannot save tag shares!")
		}
	}

	tag = models.Tag{
		ID:      id.String(),
		Name:    input.Name,
		OwnerID: ownerID,
	}

	if err := r.DB.Save(&tag).Error; err != nil {
		return &tag, gqlerror.Errorf("Incorrect form data or tag already exists!")
	}

	return &tag, nil
}

func (r *mutationResolver) TagUpdate(ctx context.Context, id string, input models.TagUpdateInput) (*models.Tag, error) {
	var tag models.Tag

	if err := r.Validator.Validate(input); err != nil {
		return &tag, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tag, err
	}

	ownerID := claims["id"].(string)

	if err := r.DB.Where("id = ? AND owner_id = ?", id, ownerID).First(&tag).Error; err != nil {
		return &tag, gqlerror.Errorf("Tag not found or you are not the owner!")
	}

	if input.Name != "" {
		tag.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if err := r.DB.Where("id = ?", input.OwnerID).First(&models.User{}).Error; err != nil {
			return nil, gqlerror.Errorf("Owner not found!")
		}

		tag.OwnerID = input.OwnerID
	}

	if len(input.SharedFor) > 0 {
		// Update tag shares
		tagShares := utils.TagShareInputsToTagShares(tag.ID, input.SharedFor)

		if err := r.DB.Save(&tagShares).Error; err != nil {
			return &tag, gqlerror.Errorf("Cannot update tag shares!")
		}
	}

	if err := r.DB.Save(&tag).Error; err != nil {
		return &tag, gqlerror.Errorf("Incorrect form data or tag already exists!")
	}

	return &tag, nil
}

func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*models.Tag, error) {
	var tag models.Tag

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tag, err
	}

	ownerID := claims["id"].(string)

	if err := r.DB.Where("id = ? AND owner_id = ?", id, ownerID).First(&tag).Delete(&tag).Error; err != nil {
		return &tag, gqlerror.Errorf("Tag not found or you are not the owner!")
	}

	// Delete tag shares
	var tagShares []models.TagShare

	if err := r.DB.Where("tag_id = ?", id).Find(&tagShares).Delete(&tagShares).Error; err != nil {
		return &tag, gqlerror.Errorf("Cannot delete tag shares!")
	}

	return &tag, nil
}

// Field resolver

func (r *tagResolver) Owner(ctx context.Context, obj *models.Tag) (*models.User, error) {
	var owner models.User

	if err := r.DB.Where("id = ?", obj.OwnerID).First(&owner).Error; err != nil {
		return &owner, gqlerror.Errorf("Owner not found!")
	}

	return &owner, nil
}

func (r *tagResolver) SharedFor(ctx context.Context, obj *models.Tag) ([]*models.TagShare, error) {
	var tagShares []*models.TagShare

	if err := r.DB.Where("tag_id = ?", obj.ID).Find(&tagShares).Error; err != nil {
		return tagShares, gqlerror.Errorf("Internal database error occurred while getting tag shares!")
	}

	return tagShares, nil
}
