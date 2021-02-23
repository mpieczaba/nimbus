package utils

import (
	"github.com/mpieczaba/nimbus/file"
)

func TagIDsToFileTags(fileId string, tagIDs []string) []file.FileTag {
	var fileTags []file.FileTag

	for _, tagID := range tagIDs {
		fileTags = append(fileTags, file.FileTag{
			FileID: fileId,
			TagID:  tagID,
		})
	}

	return fileTags
}
