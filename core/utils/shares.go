package utils

import "github.com/mpieczaba/nimbus/core/models"

func TagShareInputsToTagShares(tagID string, tagShareInputs []models.TagShareInput) []models.TagShare {
	var tagShares []models.TagShare

	for _, tagShareInput := range tagShareInputs {
		tagShares = append(tagShares, models.TagShare{
			TagID:       tagID,
			UserID:      tagShareInput.UserID,
			Permissions: tagShareInput.Permissions,
		})
	}

	return tagShares
}

func FileShareInputsToFileShares(fileID string, fileShareInputs []models.FileShareInput) []models.FileShare {
	var fileShares []models.FileShare

	for _, fileShareInput := range fileShareInputs {
		fileShares = append(fileShares, models.FileShare{
			FileID:      fileID,
			UserID:      fileShareInput.UserID,
			Permissions: fileShareInput.Permissions,
		})
	}

	return fileShares
}
