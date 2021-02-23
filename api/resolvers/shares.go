package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
)

// TagShare

// Field resolver

func (r *tagShareResolver) User(ctx context.Context, obj *models.TagShare) (*user.User, error) {
	return r.UserStore.GetUserById(obj.UserID)
}

func (r *fileShareResolver) User(ctx context.Context, obj *file.FileShare) (*user.User, error) {
	return r.UserStore.GetUserById(obj.UserID)
}
