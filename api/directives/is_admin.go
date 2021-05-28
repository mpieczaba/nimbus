package directives

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func IsAdmin() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		claims, err := auth.ClaimsFromContext(ctx)

		if err != nil {
			return nil, err
		}

		if claims.Kind != models.UserKindAdmin {
			return nil, gqlerror.Errorf("User must be an admin!")
		}

		return next(ctx)
	}
}
