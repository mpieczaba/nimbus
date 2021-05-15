package auth

import (
	"context"
	"net/http"

	"github.com/mpieczaba/nimbus/user"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}

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

		ctx := context.WithValue(c.Request.Context(), UserCtxKey, &user.User{
			ID:       claims.ID,
			Username: claims.Username,
			Kind:     claims.Kind,
		})

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
