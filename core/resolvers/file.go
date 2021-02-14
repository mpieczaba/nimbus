package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/core/utils"

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

	id := xid.New()

	// Write file in data directory
	if err := utils.WriteFile(id.String(), input.File.File); err != nil {
		return &file, gqlerror.Errorf("Cannot save file!")
	}

	file = models.File{
		ID:        id.String(),
		Name:      input.Name,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
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

	// Query file to update
	if err := r.DB.Where("id = ?", id).First(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	if input.Name != "" {
		file.Name = input.Name
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

	if err := r.DB.Where("id = ?", id).First(&file).Delete(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	// Delete file in data directory
	if err := utils.RemoveFile(id); err != nil {
		return &file, gqlerror.Errorf("Cannot delete file!")
	}

	return &file, nil
}
