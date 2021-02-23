package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*file.File, error) {
	var queryFile file.File

	if err := r.DB.Where("id = ?", id).First(&queryFile).Error; err != nil {
		return &queryFile, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	return &queryFile, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*file.File, error) {
	var files []*file.File

	if err := r.DB.Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}

// Mutation

func (r *mutationResolver) FileCreate(ctx context.Context, input file.FileInput) (*file.File, error) {
	var fileToCreate file.File

	if err := r.Validator.Validate(input); err != nil {
		return &fileToCreate, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &fileToCreate, err
	}

	id := xid.New()

	// Write file in data directory
	if err := utils.WriteFile(id.String(), input.File.File); err != nil {
		return &fileToCreate, gqlerror.Errorf("Cannot save file!")
	}

	// Save file tags
	fileTags := utils.TagIDsToFileTags(id.String(), input.Tags)

	if err := r.DB.Save(&fileTags).Error; err != nil {
		return &fileToCreate, gqlerror.Errorf("Cannot save file tags!")
	}

	if len(input.SharedFor) > 0 {
		// Save file shares
		fileShares := utils.FileShareInputsToFileShares(id.String(), input.SharedFor)

		if err := r.DB.Save(&fileShares).Error; err != nil {
			return &fileToCreate, gqlerror.Errorf("Cannot save file shares!")
		}
	}

	fileToCreate = file.File{
		ID:        id.String(),
		Name:      input.Name,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
		OwnerID:   claims["id"].(string),
	}

	if err := r.DB.Save(&fileToCreate).Error; err != nil {
		return &fileToCreate, gqlerror.Errorf("Incorrect form data!")
	}

	return &fileToCreate, nil
}

func (r *mutationResolver) FileUpdate(ctx context.Context, id string, input file.FileUpdateInput) (*file.File, error) {
	var fileToUpdate file.File

	if err := r.Validator.Validate(input); err != nil {
		return &fileToUpdate, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &fileToUpdate, err
	}

	// Query file to update
	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&fileToUpdate).Error; err != nil {
		return &fileToUpdate, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	if input.Name != "" {
		fileToUpdate.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if _, err := r.UserStore.GetUserById(input.OwnerID); err != nil {
			return nil, err
		}

		fileToUpdate.OwnerID = input.OwnerID
	}

	if input.Tags[0] != "" {
		// Update file tags
		fileTags := utils.TagIDsToFileTags(fileToUpdate.ID, input.Tags)

		if err := r.DB.Save(&fileTags).Error; err != nil {
			return &fileToUpdate, gqlerror.Errorf("Cannot update file tags!")
		}
	}

	if len(input.SharedFor) > 0 {
		// Update file shares
		fileShares := utils.FileShareInputsToFileShares(fileToUpdate.ID, input.SharedFor)

		if err := r.DB.Save(&fileShares).Error; err != nil {
			return &fileToUpdate, gqlerror.Errorf("Cannot update file shares!")
		}
	}

	if input.File.File != nil {
		// Write file in data directory
		if err := utils.WriteFile(fileToUpdate.ID, input.File.File); err != nil {
			return &fileToUpdate, gqlerror.Errorf("Cannot save file!")
		}

		fileToUpdate.MimeType = input.File.ContentType
		fileToUpdate.Extension = filepath.Ext(input.File.Filename)
		fileToUpdate.Size = input.File.Size
	}

	if err := r.DB.Save(&fileToUpdate).Error; err != nil {
		return &fileToUpdate, gqlerror.Errorf("Incorrect form data!")
	}

	return &fileToUpdate, nil
}

func (r *mutationResolver) FileDelete(ctx context.Context, id string) (*file.File, error) {
	var fileToDelete file.File

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &fileToDelete, err
	}

	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&fileToDelete).Delete(&fileToDelete).Error; err != nil {
		return &fileToDelete, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	// Delete file tags
	var fileTags []file.FileTag

	if err := r.DB.Where("file_id = ?", id).Find(&fileTags).Delete(&fileTags).Error; err != nil {
		return &fileToDelete, gqlerror.Errorf("Cannot delete file tags!")
	}

	// Delete file shares
	var fileShares []file.FileShare

	if err := r.DB.Where("file_id = ?", id).Find(&fileShares).Delete(&fileShares).Error; err != nil {
		return &fileToDelete, gqlerror.Errorf("Cannot delete file shares!")
	}

	// Delete file in data directory
	if err := utils.RemoveFile(id); err != nil {
		return &fileToDelete, gqlerror.Errorf("Cannot delete file!")
	}

	return &fileToDelete, nil
}

// Field resolver

func (r *fileResolver) Owner(ctx context.Context, obj *file.File) (*user.User, error) {
	return r.UserStore.GetUserById(obj.OwnerID)
}

func (r *fileResolver) Tags(ctx context.Context, obj *file.File) ([]*models.Tag, error) {
	var tags []*models.Tag

	tagsIDs := r.DB.Select("tag_id").Where("file_id = ?", obj.ID).Table("file_tags")

	if err := r.DB.Where("id IN (?)", tagsIDs).Find(&tags).Error; err != nil {
		return tags, gqlerror.Errorf("Internal database error occurred while getting file tags!")
	}

	return tags, nil
}

func (r *fileResolver) SharedFor(ctx context.Context, obj *file.File) ([]*file.FileShare, error) {
	var fileShares []*file.FileShare

	if err := r.DB.Where("file_id = ?", obj.ID).Find(&fileShares).Error; err != nil {
		return fileShares, gqlerror.Errorf("Internal database error occurred while getting file shares!")
	}

	return fileShares, nil
}
