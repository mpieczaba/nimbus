package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/mpieczaba/nimbus/models"
)

// Mutation

func (r *mutationResolver) AddCollaboratorToFile(ctx context.Context, input models.FileCollaboratorInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	if _, err := r.Store.FileCollaborator.CreateFileCollaborator(claims, &models.FileCollaborator{
		FileID:         input.FileID,
		CollaboratorID: input.CollaboratorID,
		Permissions:    utils.GetFilePermissionsIndex(input.Permissions),
	}); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile(claims, models.FilePermissionsRead, "id = ?", input.FileID)
}

func (r *mutationResolver) RemoveCollaboratorFromFile(ctx context.Context, fileID, collaboratorID string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	if _, err := r.Store.FileCollaborator.DeleteFileCollaborator(claims, fileID, collaboratorID); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile(claims, models.FilePermissionsRead, "id = ?", fileID)
}
