package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*models.File, error) {
	var file models.File

	if err := r.DB.Where("id = ?", id).First(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	return &file, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*models.File, error) {
	var files []*models.File

	if err := r.DB.Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}

// Mutation

func (r *mutationResolver) FileCreate(ctx context.Context, input models.FileInput) (*models.File, error) {
	var file models.File

	if err := r.Validator.Validate(input); err != nil {
		return &file, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &file, err
	}

	id := xid.New()

	// Write file in data directory
	if err := utils.WriteFile(id.String(), input.File.File); err != nil {
		return &file, gqlerror.Errorf("Cannot save file!")
	}

	// Save file tags
	fileTags := utils.TagIDsToFileTags(id.String(), input.Tags)

	if err := r.DB.Save(&fileTags).Error; err != nil {
		return &file, gqlerror.Errorf("Cannot save file tags!")
	}

	if len(input.SharedFor) > 0 {
		// Save file shares
		fileShares := utils.FileShareInputsToFileShares(id.String(), input.SharedFor)

		if err := r.DB.Save(&fileShares).Error; err != nil {
			return &file, gqlerror.Errorf("Cannot save file shares!")
		}
	}

	file = models.File{
		ID:        id.String(),
		Name:      input.Name,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
		OwnerID:   claims["id"].(string),
	}

	if err := r.DB.Save(&file).Error; err != nil {
		return &file, gqlerror.Errorf("Incorrect form data!")
	}

	return &file, nil
}

func (r *mutationResolver) FileUpdate(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	var file models.File

	if err := r.Validator.Validate(input); err != nil {
		return &file, err
	}

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &file, err
	}

	// Query file to update
	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	if input.Name != "" {
		file.Name = input.Name
	}

	if input.OwnerID != "" {
		// Check if owner does exist
		if _, err := r.UserStore.GetUserById(input.OwnerID); err != nil {
			return nil, err
		}

		file.OwnerID = input.OwnerID
	}

	if input.Tags[0] != "" {
		// Update file tags
		fileTags := utils.TagIDsToFileTags(file.ID, input.Tags)

		if err := r.DB.Save(&fileTags).Error; err != nil {
			return &file, gqlerror.Errorf("Cannot update file tags!")
		}
	}

	if len(input.SharedFor) > 0 {
		// Update file shares
		fileShares := utils.FileShareInputsToFileShares(file.ID, input.SharedFor)

		if err := r.DB.Save(&fileShares).Error; err != nil {
			return &file, gqlerror.Errorf("Cannot update file shares!")
		}
	}

	if input.File.File != nil {
		// Write file in data directory
		if err := utils.WriteFile(file.ID, input.File.File); err != nil {
			return &file, gqlerror.Errorf("Cannot save file!")
		}

		file.MimeType = input.File.ContentType
		file.Extension = filepath.Ext(input.File.Filename)
		file.Size = input.File.Size
	}

	if err := r.DB.Save(&file).Error; err != nil {
		return &file, gqlerror.Errorf("Incorrect form data!")
	}

	return &file, nil
}

func (r *mutationResolver) FileDelete(ctx context.Context, id string) (*models.File, error) {
	var file models.File

	claims, err := utils.Auth(r.Ctx)

	if err != nil {
		return &file, err
	}

	if err := r.DB.Where("id = ? AND owner_id = ?", id, claims["id"].(string)).First(&file).Delete(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	// Delete file tags
	var fileTags []models.FileTag

	if err := r.DB.Where("file_id = ?", id).Find(&fileTags).Delete(&fileTags).Error; err != nil {
		return &file, gqlerror.Errorf("Cannot delete file tags!")
	}

	// Delete file shares
	var fileShares []models.FileShare

	if err := r.DB.Where("file_id = ?", id).Find(&fileShares).Delete(&fileShares).Error; err != nil {
		return &file, gqlerror.Errorf("Cannot delete file shares!")
	}

	// Delete file in data directory
	if err := utils.RemoveFile(id); err != nil {
		return &file, gqlerror.Errorf("Cannot delete file!")
	}

	return &file, nil
}

// Field resolver

func (r *fileResolver) Owner(ctx context.Context, obj *models.File) (*user.User, error) {
	return r.UserStore.GetUserById(obj.OwnerID)
}

func (r *fileResolver) Tags(ctx context.Context, obj *models.File) ([]*models.Tag, error) {
	var tags []*models.Tag

	tagsIDs := r.DB.Select("tag_id").Where("file_id = ?", obj.ID).Table("file_tags")

	if err := r.DB.Where("id IN (?)", tagsIDs).Find(&tags).Error; err != nil {
		return tags, gqlerror.Errorf("Internal database error occurred while getting file tags!")
	}

	return tags, nil
}

func (r *fileResolver) SharedFor(ctx context.Context, obj *models.File) ([]*models.FileShare, error) {
	var fileShares []*models.FileShare

	if err := r.DB.Where("file_id = ?", obj.ID).Find(&fileShares).Error; err != nil {
		return fileShares, gqlerror.Errorf("Internal database error occurred while getting file shares!")
	}

	return fileShares, nil
}