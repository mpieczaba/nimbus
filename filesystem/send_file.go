package filesystem

import (
	"net/http"
	"os"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store"

	"github.com/gin-gonic/gin"
)

func SendFile(store *store.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := auth.ClaimsFromContext(c.Request.Context())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "User must be signed in!",
				"data":    nil,
			})

			return
		}

		if _, err = store.File.GetFile(claims, models.FilePermissionRead, "id = ?", c.Param("id")); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "File not found!",
				"data":    nil,
			})

			return
		}

		c.File(os.Getenv("DATA_DIRECTORY_PATH") + "/" + c.Param("id"))
	}
}
