package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func FileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.File(os.Getenv("DATA_DIRECTORY_PATH") + "/" + c.Param("id"))
	}
}
