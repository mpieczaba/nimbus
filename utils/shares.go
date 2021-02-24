package utils

import (
	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/file"
)

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

func FileShareInputsToFileShares(fileID string, fileShareInputs []file.FileShareInput) []*file.FileShare {
	var fileShares []*file.FileShare

	for _, fileShareInput := range fileShareInputs {
		fileShares = append(fileShares, &file.FileShare{
			FileID:      fileID,
			UserID:      fileShareInput.UserID,
			Permissions: fileShareInput.Permissions,
		})
	}

	return fileShares
}
