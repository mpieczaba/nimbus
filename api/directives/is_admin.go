package directives

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func IsAdmin() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		claims, err := auth.GetAuthClaimsFromContext(ctx)

		if err != nil {
			return nil, err
		}

		if claims.Kind != "Admin" {
			return nil, gqlerror.Errorf("User must be an admin!")
		}

		return next(ctx)
	}
}
