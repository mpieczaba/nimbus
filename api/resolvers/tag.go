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
	return r.TagStore.GetTag("id = ?", id)
}

func (r *queryResolver) Tags(ctx context.Context) ([]*tag.Tag, error) {
	return r.TagStore.GetAllTags()
}

// Mutation

func (r *mutationResolver) TagCreate(ctx context.Context, input tag.TagInput) (*tag.Tag, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	id := xid.New()

	if len(input.SharedFor) > 0 {
		// Save tag shares
		tagShares := utils.TagShareInputsToTagShares(id.String(), input.SharedFor)

		if err := r.DB.Save(&tagShares).Error; err != nil {
			return nil, gqlerror.Errorf("Cannot save tag shares!")
		}
	}

	return r.TagStore.SaveTag(&tag.Tag{
		ID:      id.String(),
		Name:    input.Name,
		OwnerID: claims["id"].(string),
	})
}

func (r *mutationResolver) TagUpdate(ctx context.Context, id string, input tag.TagUpdateInput) (*tag.Tag, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	tagToUpdate, err := r.TagStore.GetTag("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
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
			return nil, gqlerror.Errorf("Cannot update tag shares!")
		}
	}

	return r.TagStore.SaveTag(tagToUpdate)
}

func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*tag.Tag, error) {
	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	tagToDelete, err := r.TagStore.DeleteTag("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	// Delete tag shares
	var tagShares []tag.TagShare

	if err := r.DB.Where("tag_id = ?", id).Find(&tagShares).Delete(&tagShares).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete tag shares!")
	}

	return tagToDelete, nil
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
