package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

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

	id := xid.New()

	file = models.File{
		ID:   id.String(),
		Name: input.Name,
	}

	// TODO: Add file saving in data directory

	if err := r.DB.Save(&file).Error; err != nil {
		return &file, gqlerror.Errorf("Incorrect form data!")
	}

	return &file, nil
}

func (r *mutationResolver) FileUpdate(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	var file models.File

	// Query file to update
	if err := r.DB.Where("id = ?", id).First(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	if input.Name != "" {
		file.Name = input.Name
	}

	// TODO: Add file updating in data directory

	if err := r.DB.Save(&file).Error; err != nil {
		return &file, gqlerror.Errorf("Incorrect form data!")
	}

	return &file, nil
}

func (r *mutationResolver) FileDelete(ctx context.Context, id string) (*models.File, error) {
	var file models.File

	// TODO: Add file deleting in data directory

	if err := r.DB.Where("id = ?", id).First(&file).Delete(&file).Error; err != nil {
		return &file, gqlerror.Errorf("File with id `" + id + "` not found!")
	}

	return &file, nil
}
