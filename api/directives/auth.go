package directives

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if ctx.Value(auth.UserCtxKey) == nil {
			return nil, gqlerror.Errorf("User must be signed in!")
		}

		return next(ctx)
	}
}
