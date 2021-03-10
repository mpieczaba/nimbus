package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*file.File, error) {
	return r.Store.File.GetFile("id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context) ([]*file.File, error) {
	return r.Store.File.GetAllFiles()
}

// Mutation

func (r *mutationResolver) FileCreate(ctx context.Context, input file.FileInput) (*file.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	id := xid.New()

	// Write file in data directory
	if err = r.Filesystem.WriteFile(id.String(), input.File.File); err != nil {
		return nil, gqlerror.Errorf("Cannot save file!")
	}

	return r.Store.File.CreateFile(&file.File{
		ID:         id.String(),
		Name:       input.Name,
		MimeType:   input.File.ContentType,
		Extension:  filepath.Ext(input.File.Filename),
		Size:       input.File.Size,
		OwnerID:    claims["id"].(string),
		FileTags:   utils.TagIDsToFileTags(input.Tags),
		FileShares: utils.FileShareInputsToFileShares(input.SharedFor),
	})
}

// TODO: Fix for editor
func (r *mutationResolver) FileUpdate(ctx context.Context, id string, input file.FileUpdateInput) (*file.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileToUpdate, err := r.Store.File.GetFile("id = ? AND owner_id = ?", id, claims["id"].(string))

	// Get file to update if user is co-owner
	if err != nil {
		query := "file_id = ? AND user_id = ? AND permissions = ?"

		fileToUpdate, err = r.Store.File.GetFile("id IN (?)", r.Store.FileShare.GetFileShareAsSubQuery(query, id, claims["id"].(string), "CoOwner"))

		if err != nil {
			return nil, err
		}
	}

	if input.Name != "" {
		fileToUpdate.Name = input.Name
	}

	if input.OwnerID != "" {
		fileToUpdate.OwnerID = input.OwnerID
	}

	if len(input.Tags) > 0 {
		// Update file tags
		fileToUpdate.FileTags = utils.TagIDsToFileTags(input.Tags)
	}

	if len(input.SharedFor) > 0 {
		// Update file shares
		fileToUpdate.FileShares = utils.FileShareInputsToFileShares(input.SharedFor)
	}

	if input.File.File != nil {
		// Write file in data directory
		if err = r.Filesystem.WriteFile(fileToUpdate.ID, input.File.File); err != nil {
			return nil, gqlerror.Errorf("Cannot save file!")
		}

		fileToUpdate.MimeType = input.File.ContentType
		fileToUpdate.Extension = filepath.Ext(input.File.Filename)
		fileToUpdate.Size = input.File.Size
	}

	return r.Store.File.UpdateFile(fileToUpdate)
}

func (r *mutationResolver) FileDelete(ctx context.Context, id string) (*file.File, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	fileToDelete, err := r.Store.File.DeleteFile("id = ? AND owner_id = ?", id, claims["id"].(string))

	// Get file to delete if user is co-owner
	if err != nil {
		query := "file_id = ? AND user_id = ? AND permissions = ?"

		fileToDelete, err = r.Store.File.DeleteFile("id = ?", r.Store.FileShare.GetFileShareAsSubQuery(query, id, claims["id"].(string), "CoOwner"))

		if err != nil {
			return nil, err
		}
	}

	// Delete file in data directory
	if err = r.Filesystem.RemoveFile(id); err != nil {
		return nil, gqlerror.Errorf("Cannot delete file!")
	}

	return fileToDelete, nil
}

// Field resolver

func (r *fileResolver) Owner(ctx context.Context, obj *file.File) (*user.User, error) {
	return r.Store.User.GetUser(obj.OwnerID)
}

func (r *fileResolver) Tags(ctx context.Context, obj *file.File) ([]*tag.Tag, error) {
	return r.Store.Tag.GetAllTagsWithCondition("id IN (?)", r.Store.FileTag.GetTagIDs("file_id = ?", obj.ID))
}

func (r *fileResolver) SharedFor(ctx context.Context, obj *file.File) ([]*file_share.FileShare, error) {
	return r.Store.FileShare.GetAllFileShares("file_id = ?", obj.ID)
}
