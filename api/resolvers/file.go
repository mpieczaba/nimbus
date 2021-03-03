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

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	id := xid.New()

	// Write file in data directory
	if err = r.Filesystem.WriteFile(id.String(), input.File.File); err != nil {
		return nil, gqlerror.Errorf("Cannot save file!")
	}

	// Save file tags
	fileTags := utils.TagIDsToFileTags(id.String(), input.Tags)

	if _, err = r.Store.File.SaveFileTags(fileTags); err != nil {
		return nil, err
	}

	return r.Store.File.SaveFile(&file.File{
		ID:        id.String(),
		Name:      input.Name,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
		OwnerID:   claims["id"].(string),
	})
}

func (r *mutationResolver) FileUpdate(ctx context.Context, id string, input file.FileUpdateInput) (*file.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	fileToUpdate, err := r.Store.File.GetFile("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		fileToUpdate.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if _, err = r.Store.User.GetUser(input.OwnerID); err != nil {
			return nil, err
		}

		fileToUpdate.OwnerID = input.OwnerID
	}

	if len(input.Tags) > 0 {
		// Update file tags
		fileTags := utils.TagIDsToFileTags(fileToUpdate.ID, input.Tags)

		if _, err = r.Store.File.SaveFileTags(fileTags); err != nil {
			return nil, err
		}
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

	return r.Store.File.SaveFile(fileToUpdate)
}

func (r *mutationResolver) FileDelete(ctx context.Context, id string) (*file.File, error) {
	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return nil, err
	}

	fileToDelete, err := r.Store.File.DeleteFile("id = ? AND owner_id = ?", id, claims["id"].(string))

	if err != nil {
		return nil, err
	}

	// Delete file tags
	if _, err = r.Store.File.DeleteFileTags("file_id = ?", id); err != nil {
		return nil, err
	}

	// Delete file shares
	if _, err = r.Store.FileShare.DeleteFileShares("file_id = ?", id); err != nil {
		return nil, err
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
	return r.Store.Tag.GetAllTagsWithCondition("id IN (?)", r.Store.File.GetTagIDs("file_id = ?", obj.ID))
}

func (r *fileResolver) SharedFor(ctx context.Context, obj *file.File) ([]*file_share.FileShare, error) {
	return r.Store.FileShare.GetAllFileShares("file_id = ?", obj.ID)
}
