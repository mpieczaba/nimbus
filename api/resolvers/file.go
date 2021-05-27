package resolvers

import (
	"context"
	"path/filepath"

	"github.com/mpieczaba/nimbus/file"

	"github.com/rs/xid"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*file.File, error) {
	return r.Store.File.GetFile("id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int) (*file.FileConnection, error) {
	return r.Store.File.GetAllFiles(after, before, first, last)
}

// Mutation

func (r *mutationResolver) CreateFile(ctx context.Context, input file.FileInput) (*file.File, error) {
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

	return r.Store.File.CreateFile(&file.File{
		ID:        xid.New().String(),
		Name:      fileName,
		MimeType:  input.File.ContentType,
		Extension: filepath.Ext(input.File.Filename),
		Size:      input.File.Size,
	})
}
