package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/models"
)

// Mutation

func (r *mutationResolver) AddFileCollaborator(ctx context.Context, input models.FileCollaboratorInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	if _, err := r.Store.FileCollaborator.AddFileCollaborator(&models.FileCollaborator{
		FileID:         input.FileID,
		CollaboratorID: input.CollaboratorID,
		Permission:     input.Permission,
	}); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile("id = ?", input.FileID)
}
