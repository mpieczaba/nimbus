package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/tag/tag_share"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
)

// Mutation

func (r *mutationResolver) TagShareCreate(ctx context.Context, input tag_share.TagShareInput) (*tag_share.TagShare, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.Tag.GetTag("id = ? AND owner_id = ?", input.TagID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.TagShare.SaveTagShare(&tag_share.TagShare{
		ID:          xid.New().String(),
		TagID:       input.TagID,
		UserID:      input.UserID,
		Permissions: input.Permissions,
	})
}

func (r *mutationResolver) TagShareUpdate(ctx context.Context, id string, input tag_share.TagShareUpdateInput) (*tag_share.TagShare, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	tagShareToUpdate, err := r.Store.TagShare.GetTagShare("id = ?", id)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.Tag.GetTag("id = ? AND owner_id = ?", tagShareToUpdate.TagID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	if input.Permissions != 0 {
		tagShareToUpdate.Permissions = input.Permissions
	}

	return r.Store.TagShare.SaveTagShare(tagShareToUpdate)
}

func (r *mutationResolver) TagShareDelete(ctx context.Context, id string) (*tag_share.TagShare, error) {
	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	tagShareToDelete, err := r.Store.TagShare.GetTagShare("id = ?", id)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.Tag.GetTag("id = ? AND owner_id = ?", tagShareToDelete.TagID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.TagShare.DeleteTagShare(id)
}

// Field resolver

func (r *tagShareResolver) Tag(ctx context.Context, obj *tag_share.TagShare) (*tag.Tag, error) {
	return r.Store.Tag.GetTag("id = ?", obj.TagID)
}

func (r *tagShareResolver) User(ctx context.Context, obj *tag_share.TagShare) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.UserID)
}
