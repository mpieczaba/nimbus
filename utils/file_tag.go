package utils

import (
	"github.com/mpieczaba/nimbus/file/file_tag"
)

func TagIDsToFileTags(tagIDs []string) []file_tag.FileTag {
	var fileTags []file_tag.FileTag

	for _, tagID := range tagIDs {
		fileTags = append(fileTags, file_tag.FileTag{
			TagID: tagID,
		})
	}

	return fileTags
}
