package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
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

		if _, err = r.TagStore.SaveTagShares(tagShares); err != nil {
			return nil, err
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
		if _, err = r.UserStore.GetUser("id = ?", input.OwnerID); err != nil {
			return nil, err
		}

		tagToUpdate.OwnerID = input.OwnerID
	}

	if len(input.SharedFor) > 0 {
		// Update tag shares
		tagShares := utils.TagShareInputsToTagShares(tagToUpdate.ID, input.SharedFor)

		if _, err = r.TagStore.SaveTagShares(tagShares); err != nil {
			return nil, err
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
	if _, err = r.TagStore.DeleteTagShares("tag_id = ?", id); err != nil {
		return nil, err
	}

	return tagToDelete, nil
}

// Field resolver

func (r *tagResolver) Owner(ctx context.Context, obj *tag.Tag) (*user.User, error) {
	return r.UserStore.GetUser("id = ?", obj.OwnerID)
}

func (r *tagResolver) SharedFor(ctx context.Context, obj *tag.Tag) ([]*tag.TagShare, error) {
	return r.TagStore.GetAllTagShares("tag_id = ?", obj.ID)
}
