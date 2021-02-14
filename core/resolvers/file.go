package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"

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
