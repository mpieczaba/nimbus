package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
)

// TagShare

// Field resolver

func (r *fileShareResolver) User(ctx context.Context, obj *file.FileShare) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.UserID)
}
