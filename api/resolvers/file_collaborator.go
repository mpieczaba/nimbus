package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Mutation

func (r *mutationResolver) AddFileCollaborator(ctx context.Context, input models.FileCollaboratorInput) (*models.File, error) {

	return nil, gqlerror.Errorf("Not implemented!")
}
