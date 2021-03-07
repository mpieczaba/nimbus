package utils

import (
	"github.com/mpieczaba/nimbus/file"
)

func TagIDsToFileTags(tagIDs []string) []file.FileTag {
	var fileTags []file.FileTag

	for _, tagID := range tagIDs {
		fileTags = append(fileTags, file.FileTag{
			TagID: tagID,
		})
	}

	return fileTags
}
