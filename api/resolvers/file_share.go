package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/user"
)

// Mutation

func (r *mutationResolver) FileShareDelete(ctx context.Context, fileId string, userId string) (*file_share.FileShare, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileShareToDelete, err := r.Store.FileShare.GetFileShare("file_id = ? AND user_id = ?", fileId, userId)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND owner_id = ?", fileShareToDelete.FileID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.FileShare.DeleteFileShare("file_id = ? AND user_id = ?", fileId, userId)
}

// Field resolver

func (r *fileShareResolver) User(ctx context.Context, obj *file_share.FileShare) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.UserID)
}
