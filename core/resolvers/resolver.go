package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/mpieczaba/nimbus/core/generated"
)

type Resolver struct{}

func (r *queryResolver) Test(ctx context.Context) (*string, error) {
	panic("not implemented")
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
