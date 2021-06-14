package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
)

// Query

func (r *queryResolver) Tag(ctx context.Context, name string) (*models.Tag, error) {
	return r.Store.Tag.GetTag("name = ?", name)
}

func (r *queryResolver) Tags(ctx context.Context, after *string, before *string, first *int, last *int, name *string) (*models.TagConnection, error) {
	return r.Store.Tag.GetAllTags(after, before, first, last, name)
}

// Field resolver

func (r *tagResolver) Files(ctx context.Context, obj *models.Tag, after, before *string, first, last *int, name *string, permission *models.FilePermission, tags []string) (*models.FileConnection, error) {
	claims, _ := auth.ClaimsFromContext(ctx)

	return r.Store.File.GetAllFiles(claims, after, before, first, last, name, *permission, append(tags, obj.Name))
}
