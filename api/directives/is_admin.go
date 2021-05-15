package directives

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/user"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func IsAdmin() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if ctx.Value(auth.UserCtxKey) == nil {
			return nil, gqlerror.Errorf("User must be signed in!")
		}

		if ctx.Value(auth.UserCtxKey).(*user.User).Kind != "Admin" {
			return nil, gqlerror.Errorf("User must be an admin!")
		}

		return next(ctx)
	}
}
