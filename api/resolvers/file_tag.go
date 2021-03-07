package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file/file_tag"
)

// Mutation

func (r *mutationResolver) FileTagDelete(ctx context.Context, fileId string, tagId string) (*file_tag.FileTag, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileTagToDelete, err := r.Store.FileTag.GetFileTag("file_id = ? AND tag_id = ?", fileId, tagId)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND owner_id = ?", fileTagToDelete.FileID, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.FileTag.DeleteFileTag("file_id = ? AND tag_id = ?", fileId, tagId)
}
