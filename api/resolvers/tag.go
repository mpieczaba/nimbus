package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/tag/tag_share"
	"github.com/mpieczaba/nimbus/user"

	"github.com/rs/xid"
)

// Query

func (r *queryResolver) Tag(ctx context.Context, id string) (*tag.Tag, error) {
	return r.Store.Tag.GetTag("id = ?", id)
}

func (r *queryResolver) Tags(ctx context.Context) ([]*tag.Tag, error) {
	return r.Store.Tag.GetAllTags()
}

// Mutation

func (r *mutationResolver) TagCreate(ctx context.Context, input tag.TagInput) (*tag.Tag, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	return r.Store.Tag.CreateTag(&tag.Tag{
		ID:      xid.New().String(),
		Name:    input.Name,
		OwnerID: claims["id"].(string),
	})
}

func (r *mutationResolver) TagUpdate(ctx context.Context, id string, input tag.TagUpdateInput) (*tag.Tag, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	tagToUpdate, err := r.Store.Tag.GetTag("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		tagToUpdate.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if _, err = r.Store.User.GetUser("id = ?", input.OwnerID); err != nil {
			return nil, err
		}

		tagToUpdate.OwnerID = input.OwnerID
	}

	return r.Store.Tag.UpdateTag(tagToUpdate)
}

func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*tag.Tag, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	tagToDelete, err := r.Store.Tag.DeleteTag("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	// Delete tag shares
	if _, err = r.Store.TagShare.DeleteTagShares("tag_id = ?", id); err != nil {
		return nil, err
	}

	return tagToDelete, nil
}

// Field resolver

func (r *tagResolver) Owner(ctx context.Context, obj *tag.Tag) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.OwnerID)
}

func (r *tagResolver) SharedFor(ctx context.Context, obj *tag.Tag) ([]*tag_share.TagShare, error) {
	return r.Store.TagShare.GetAllTagShares("tag_id = ?", obj.ID)
}
