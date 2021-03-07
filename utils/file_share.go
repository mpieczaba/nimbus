package utils

import (
	"github.com/mpieczaba/nimbus/file/file_share"
)

func FileShareInputsToFileShares(fileShareInputs []file_share.FileShareInput) []file_share.FileShare {
	var fileShares []file_share.FileShare

	for _, fileShareInput := range fileShareInputs {
		fileShares = append(fileShares, file_share.FileShare{
			UserID:      fileShareInput.UserID,
			Permissions: fileShareInput.Permissions,
		})
	}

	return fileShares
}
