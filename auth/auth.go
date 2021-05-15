package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if !IsToken(token) {
			c.Next()

			return
		}

		claims, err := CheckToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "User must be signed in!",
				"data":    nil,
			})

			return
		}

		ctx := context.WithValue(c.Request.Context(), userCtxKey, claims)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GetAuthClaimsFromContext(ctx context.Context) (*Claims, error) {
	claims, ok := ctx.Value(userCtxKey).(*Claims)

	if !ok {
		return nil, gqlerror.Errorf("User must be signed in!")
	}

	return claims, nil
}
