package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file/file_tag"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Mutation

func (r *mutationResolver) FileTagDelete(ctx context.Context, fileId string, tagId string) (*file_tag.FileTag, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	// Check if user is banned
	if claims["kind"].(string) == "Banned" {
		return nil, gqlerror.Errorf("You have no permissions to delete file tag!")
	}

	fileTagToDelete, err := r.Store.FileTag.GetFileTag("file_id = ? AND tag_id = ?", fileId, tagId)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.File.GetFile("id = ? AND (owner_id = ? OR ? = 'Admin')", fileTagToDelete.FileID, claims["id"].(string), claims["kind"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.FileTag.DeleteFileTag("file_id = ? AND tag_id = ?", fileId, tagId)
}
