package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	return r.Store.File.GetFile(claims.ID, models.FilePermissionRead, "id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int, permission *models.FilePermission, collaboratorID *string) (*models.FileConnection, error) {
	var filesCollaboratorID string

	if collaboratorID != nil {
		filesCollaboratorID = *collaboratorID
	} else {
		claims, _ := auth.ClaimsFromContext(ctx)

		filesCollaboratorID = claims.ID
	}

	var filesPermission models.FilePermission

	if permission != nil {
		filesPermission = *permission
	} else {
		filesPermission = models.FilePermissionRead
	}

	return r.Store.File.GetAllFiles(after, before, first, last, filesCollaboratorID, filesPermission)
}

// Mutation

func (r *mutationResolver) CreateFile(ctx context.Context, input models.FileInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	var fileName string

	if input.Name != "" {
		fileName = input.Name
	} else {
		fileName = input.File.Filename
	}

	// TODO: Add writing file to data directory

	/*
		fileContent, err := io.ReadAll(input.File.File)

		if err != nil {
			return nil, gqlerror.Errorf("Cannot open file!")
		}
	*/

	claims, _ := auth.ClaimsFromContext(ctx)

	id := xid.New().String()

	return r.Store.File.CreateFile(&models.File{
		ID:        id,
		Name:      fileName,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
		Collaborators: []models.FileCollaborator{{
			FileID:         id,
			CollaboratorID: claims.ID,
			Permission:     utils.GetFilePermissionIndex(models.FilePermissionAdmin),
		}},
	})
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	fileToUpdate, err := r.Store.File.GetFile(claims.ID, models.FilePermissionMaintain, "id = ?", id)

	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		fileToUpdate.Name = input.Name
	}

	if input.File.File != nil {
		fileToUpdate.MimeType = input.File.ContentType
		fileToUpdate.Extension = filepath.Ext(input.File.Filename)
		fileToUpdate.Size = input.File.Size
	}

	return r.Store.File.UpdateFile(fileToUpdate)
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	fileToDelete, err := r.Store.File.GetFile(claims.ID, models.FilePermissionAdmin, "id = ?", id)

	if err != nil {
		return nil, err
	}

	return r.Store.File.DeleteFile(fileToDelete)
}

// Field resolver

func (r *fileResolver) Collaborators(ctx context.Context, obj *models.File, after, before *string, first, last *int, permission *models.FilePermission) (*models.FileCollaboratorConnection, error) {
	var filesPermission models.FilePermission

	if permission != nil {
		filesPermission = *permission
	} else {
		filesPermission = models.FilePermissionRead
	}

	return r.Store.FileCollaborator.GetFileCollaborators(after, before, first, last, obj.ID, filesPermission)
}
