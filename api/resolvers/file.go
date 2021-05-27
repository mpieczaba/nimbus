package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file"
)

// Query

func (r *queryResolver) File(ctx context.Context, id string) (*file.File, error) {
	return r.Store.File.GetFile("id = ?", id)
}

func (r *queryResolver) Files(ctx context.Context, after, before *string, first, last *int) (*file.FileConnection, error) {
	return r.Store.File.GetAllFiles(after, before, first, last)
}
