package auth

import (
	"context"

	"github.com/mpieczaba/nimbus/user"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type contextKey struct {
	name string
}

var tokenCtxKey = &contextKey{"token"}
var UserCtxKey = &contextKey{"user"}

// Middleware that passes authorization token to the @auth directive
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), tokenCtxKey, c.GetHeader("Authorization"))

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func Directive() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		claims, err := CheckToken(ctx.Value(tokenCtxKey).(string))

		if err != nil {
			return nil, gqlerror.Errorf("User must be signed in!")
		}

		ctx = context.WithValue(ctx, UserCtxKey, &user.User{
			ID:       claims.ID,
			Username: claims.Username,
			Kind:     claims.Kind,
		})

		return next(ctx)
	}
}
