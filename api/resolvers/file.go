package resolvers

import (
	"context"
	"os"
	"path/filepath"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/filesystem"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/rs/xid"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	return r.Store.File.GetFile(claims, models.FilePermissionsRead, "id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int, name *string, permissions *models.FilePermissions, tags []string) (*models.FileConnection, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	return r.Store.File.GetAllFiles(claims, after, before, first, last, name, *permissions, tags)
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
			Permissions:    utils.GetFilePermissionsIndex(models.FilePermissionsAdmin),
		}},
	}, func() error {
		// Write file to data directory
		return filesystem.WriteFile(id, input.File.File)
	})
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	fileToUpdate, err := r.Store.File.GetFile(claims, models.FilePermissionsMaintain, "id = ?", id)

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

	return r.Store.File.UpdateFile(fileToUpdate, func() error {
		// Update file in data directory
		return filesystem.WriteFile(fileToUpdate.ID, input.File.File)
	})
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	fileToDelete, err := r.Store.File.GetFile(claims, models.FilePermissionsAdmin, "id = ?", id)

	if err != nil {
		return nil, err
	}

	return r.Store.File.DeleteFile(fileToDelete, func() error {
		// Remove file from data directory
		return filesystem.RemoveFile(fileToDelete.ID)
	})
}

// Field resolver

func (r *fileResolver) URL(ctx context.Context, obj *models.File) (string, error) {
	return "http://" + os.Getenv("HOST") + "/files/" + obj.ID, nil
}

func (r *fileResolver) DownloadURL(ctx context.Context, obj *models.File) (string, error) {
	return "http://" + os.Getenv("HOST") + "/files/download/" + obj.ID + "/" + obj.Name, nil
}

func (r *fileResolver) Tags(ctx context.Context, obj *models.File, after, before *string, first, last *int, name *string) (*models.FileTagConnection, error) {
	return r.Store.FileTag.GetFileTags(after, before, first, last, obj.ID, name)
}

func (r *fileResolver) Collaborators(ctx context.Context, obj *models.File, after, before *string, first, last *int, username *string, permissions *models.FilePermissions) (*models.FileCollaboratorConnection, error) {
	return r.Store.FileCollaborator.GetFileCollaborators(after, before, first, last, obj.ID, username, *permissions)
}
