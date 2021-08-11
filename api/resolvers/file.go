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

	return r.Store.File.GetFile(claims, models.FilePermissionRead, "id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int, name *string, permission *models.FilePermission, tags []string) (*models.FileConnection, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	return r.Store.File.GetAllFiles(claims, after, before, first, last, name, *permission, tags)
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

	tx, file, err := r.Store.File.CreateFile(&models.File{
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

	if err != nil {
		return nil, err
	}

	// Write file to data directory
	if err = filesystem.WriteFile(id, input.File.File); err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return file, nil
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	claims, _ := auth.ClaimsFromContext(ctx)

	fileToUpdate, err := r.Store.File.GetFile(claims, models.FilePermissionMaintain, "id = ?", id)

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

	tx, file, err := r.Store.File.UpdateFile(fileToUpdate)

	if err != nil {
		return nil, err
	}

	if input.File.File != nil {
		// Update file in data directory
		if err = filesystem.WriteFile(fileToUpdate.ID, input.File.File); err != nil {
			tx.Rollback()

			return nil, err
		}
	}

	tx.Commit()

	return file, err
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*models.File, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	fileToDelete, err := r.Store.File.GetFile(claims, models.FilePermissionAdmin, "id = ?", id)

	if err != nil {
		return nil, err
	}

	tx, file, err := r.Store.File.DeleteFile(fileToDelete)

	if err != nil {
		return nil, err
	}

	// Remove file from data directory
	if err = filesystem.RemoveFile(fileToDelete.ID); err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return file, err
}

// Field resolver

func (r *Resolver) URL(ctx context.Context, obj *models.File) (string, error) {
	return "http://" + os.Getenv("HOST") + "/files/" + obj.ID, nil
}

func (r *fileResolver) Tags(ctx context.Context, obj *models.File, after, before *string, first, last *int, name *string) (*models.FileTagConnection, error) {
	return r.Store.FileTag.GetFileTags(after, before, first, last, obj.ID, name)
}

func (r *fileResolver) Collaborators(ctx context.Context, obj *models.File, after, before *string, first, last *int, username *string, permission *models.FilePermission) (*models.FileCollaboratorConnection, error) {
	return r.Store.FileCollaborator.GetFileCollaborators(after, before, first, last, obj.ID, username, *permission)
}
