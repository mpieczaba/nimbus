package directives

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"

	"github.com/99designs/gqlgen/graphql"
)

func Auth() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if _, err := auth.GetAuthClaimsFromContext(ctx); err != nil {
			return nil, err
		}

		return next(ctx)
	}
}
