package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/models"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*models.File, error) {
	return r.Store.File.GetFile("id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int) (*models.FileConnection, error) {
	return r.Store.File.GetAllFiles(after, before, first, last)
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

	return r.Store.File.CreateFile(&models.File{
		ID:        xid.New().String(),
		Name:      fileName,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
	})
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id string, input models.FileUpdateInput) (*models.File, error) {
	if err := r.Validator.Validate(input); err != nil {
		return nil, err
	}

	fileToUpdate, err := r.Store.File.GetFile("id = ?", id)

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
	return r.Store.File.DeleteFile("id = ?", id)
}

// Field resolver

func (r *fileResolver) Collaborators(ctx context.Context, obj *models.File) (*models.FileCollaboratorConnection, error) {
	return nil, gqlerror.Errorf("Not implemented!")
}
