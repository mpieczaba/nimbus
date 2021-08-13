package utils

import "github.com/mpieczaba/nimbus/models"

func GetFilePermissionsIndex(permission models.FilePermissions) int8 {
	for i := range models.AllFilePermissions {
		if permission == models.AllFilePermissions[i] {
			return int8(i)
		}
	}

	return -1
}
