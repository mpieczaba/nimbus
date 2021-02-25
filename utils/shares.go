package utils

import (
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/tag"
)

func TagShareInputsToTagShares(tagID string, tagShareInputs []tag.TagShareInput) []*tag.TagShare {
	var tagShares []*tag.TagShare

	for _, tagShareInput := range tagShareInputs {
		tagShares = append(tagShares, &tag.TagShare{
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
