package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"
)

// Mutation

func (r *mutationResolver) AddTagsToFile(ctx context.Context, input models.FileTagsInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	if _, err := r.Store.Tag.CreateTagsOrAppendFileTags(claims, utils.FileTagsInputToTags(input)); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile(claims, models.FilePermissionsRead, "id = ?", input.FileID)
}

func (r *mutationResolver) RemoveTagsFromFile(ctx context.Context, input models.FileTagsInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	if _, err := r.Store.FileTag.DeleteFileTags(claims, input); err != nil {
		return nil, err
	}

	return r.Store.File.GetFile(claims, models.FilePermissionsRead, "id = ?", input.FileID)
}
