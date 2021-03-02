package utils

import (
	"github.com/mpieczaba/nimbus/file"
)

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
