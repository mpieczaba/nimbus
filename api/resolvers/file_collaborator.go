package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/mpieczaba/nimbus/models"
)

// Mutation

func (r *mutationResolver) AddFileCollaborator(ctx context.Context, input models.FileCollaboratorInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	if _, err := r.Store.FileCollaborator.AddFileCollaborator(claims, &models.FileCollaborator{
		FileID:         input.FileID,
		CollaboratorID: input.CollaboratorID,
		Permission:     utils.GetFilePermissionIndex(input.Permission),
	}); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile(claims, models.FilePermissionRead, "id = ?", input.FileID)
}
