package scopes

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"

	"gorm.io/gorm"
)

func FilePermission(model interface{}, fieldToSelect string, permission models.FilePermission, query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subQuery := db.Model(models.FileCollaborator{}).Select(fieldToSelect).Where(query, args...).Where("permission <= ?", utils.GetFilePermissionIndex(permission))

		return db.Session(&gorm.Session{NewDB: true}).Model(model).Where("id IN (?)", subQuery)
	}
}
