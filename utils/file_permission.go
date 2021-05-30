package utils

import "github.com/mpieczaba/nimbus/models"

func GetFilePermissionIndex(permission models.FilePermission) int8 {
	for i := range models.AllFilePermission {
		if permission == models.AllFilePermission[i] {
			return int8(i)
		}
	}

	return -1
}
