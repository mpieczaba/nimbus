package utils

import "github.com/mpieczaba/nimbus/core/models"

func TagIDsToFileTags(fileId string, tagIDs []string) []models.FileTag {
	var fileTags []models.FileTag

	for _, tagID := range tagIDs {
		fileTags = append(fileTags, models.FileTag{
			FileID: fileId,
			TagID:  tagID,
		})
	}

	return fileTags
}
