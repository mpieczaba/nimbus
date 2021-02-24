package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) Tag(ctx context.Context, id string) (*tag.Tag, error) {
	var tagToQuery tag.Tag

	if err := r.DB.Where("id = ?", id).First(&tagToQuery).Error; err != nil {
		return &tagToQuery, gqlerror.Errorf("Tag with id `" + id + "` not found!")
	}

	return &tagToQuery, nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*tag.Tag, error) {
	var tags []*tag.Tag

	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all tags!")
	}

	return tags, nil
}

// Mutation

func (r *mutationResolver) TagCreate(ctx context.Context, input tag.TagInput) (*tag.Tag, error) {
	var tagToCreate tag.Tag

	if err := r.Validator.Validate(input); err != nil {
		return &tagToCreate, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tagToCreate, err
	}

	id := xid.New()

	if len(input.SharedFor) > 0 {
		// Save tag shares
		tagShares := utils.TagShareInputsToTagShares(id.String(), input.SharedFor)

		if err := r.DB.Save(&tagShares).Error; err != nil {
			return &tagToCreate, gqlerror.Errorf("Cannot save tag shares!")
		}
	}

	tagToCreate = tag.Tag{
		ID:      id.String(),
		Name:    input.Name,
		OwnerID: claims["id"].(string),
	}

	if err := r.DB.Save(&tagToCreate).Error; err != nil {
		return &tagToCreate, gqlerror.Errorf("Incorrect form data or tag already exists!")
	}

	return &tagToCreate, nil
}

func (r *mutationResolver) TagUpdate(ctx context.Context, id string, input tag.TagUpdateInput) (*tag.Tag, error) {
	var tagToUpdate tag.Tag

	if err := r.Validator.Validate(input); err != nil {
		return &tagToUpdate, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tagToUpdate, err
	}

	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&tagToUpdate).Error; err != nil {
		return &tagToUpdate, gqlerror.Errorf("Tag not found or you are not the owner!")
	}

	if input.Name != "" {
		tagToUpdate.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if _, err := r.UserStore.GetUser("id = ?", input.OwnerID); err != nil {
			return nil, err
		}

		tagToUpdate.OwnerID = input.OwnerID
	}

	if len(input.SharedFor) > 0 {
		// Update tag shares
		tagShares := utils.TagShareInputsToTagShares(tagToUpdate.ID, input.SharedFor)

		if err := r.DB.Save(&tagShares).Error; err != nil {
			return &tagToUpdate, gqlerror.Errorf("Cannot update tag shares!")
		}
	}

	if err := r.DB.Save(&tagToUpdate).Error; err != nil {
		return &tagToUpdate, gqlerror.Errorf("Incorrect form data or tag already exists!")
	}

	return &tagToUpdate, nil
}

func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*tag.Tag, error) {
	var tagToDelete tag.Tag

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &tagToDelete, err
	}

	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&tagToDelete).Delete(&tagToDelete).Error; err != nil {
		return &tagToDelete, gqlerror.Errorf("Tag not found or you are not the owner!")
	}

	// Delete tag shares
	var tagShares []tag.TagShare

	if err := r.DB.Where("tag_id = ?", id).Find(&tagShares).Delete(&tagShares).Error; err != nil {
		return &tagToDelete, gqlerror.Errorf("Cannot delete tag shares!")
	}

	return &tagToDelete, nil
}

// Field resolver

func (r *tagResolver) Owner(ctx context.Context, obj *tag.Tag) (*user.User, error) {
	return r.UserStore.GetUser("id = ?", obj.OwnerID)
}

func (r *tagResolver) SharedFor(ctx context.Context, obj *tag.Tag) ([]*tag.TagShare, error) {
	var tagShares []*tag.TagShare

	if err := r.DB.Where("tag_id = ?", obj.ID).Find(&tagShares).Error; err != nil {
		return tagShares, gqlerror.Errorf("Internal database error occurred while getting tag shares!")
	}

	return tagShares, nil
}
