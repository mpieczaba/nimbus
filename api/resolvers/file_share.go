package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/user"

	"github.com/rs/xid"
)

// Mutation

func (r *mutationResolver) FileShareCreate(ctx context.Context, input file_share.FileShareInput) (*file_share.FileShare, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND owner_id = ?", input.FileID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.FileShare.SaveFileShare(&file_share.FileShare{
		ID:          xid.New().String(),
		FileID:      input.FileID,
		UserID:      input.UserID,
		Permissions: input.Permissions,
	})
}

func (r *mutationResolver) FileShareUpdate(ctx context.Context, id string, input file_share.FileShareUpdateInput) (*file_share.FileShare, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileShareToUpdate, err := r.Store.FileShare.GetFileShare("id = ?", id)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND owner_id = ?", fileShareToUpdate.FileID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	if input.Permissions != 0 {
		fileShareToUpdate.Permissions = input.Permissions
	}

	return r.Store.FileShare.SaveFileShare(fileShareToUpdate)
}

func (r *mutationResolver) FileShareDelete(ctx context.Context, id string) (*file_share.FileShare, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileShareToDelete, err := r.Store.FileShare.GetFileShare("id = ?", id)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND owner_id = ?", fileShareToDelete.FileID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.FileShare.DeleteFileShare(id)
}

// Field resolver

func (r *fileShareResolver) File(ctx context.Context, obj *file_share.FileShare) (*file.File, error) {
	return r.Store.File.GetFile("id = ?", obj.FileID)
}

func (r *fileShareResolver) User(ctx context.Context, obj *file_share.FileShare) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.UserID)
}
