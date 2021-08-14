package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func FileHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.File(os.Getenv("DATA_DIRECTORY_PATH") + "/" + c.Param("id"))
	}
}

func FileDownloadHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.FileAttachment(os.Getenv("DATA_DIRECTORY_PATH")+"/"+c.Param("id"), c.Param("name"))
	}
}
