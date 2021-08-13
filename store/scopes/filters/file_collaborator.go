package filters

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"

	"gorm.io/gorm"
)

func FilterFileCollaboratorsByFilePermissions(permissions models.FilePermissions, query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subQuery := db.Session(&gorm.Session{NewDB: true}).Model(models.FileCollaborator{}).Select("collaborator_id").Where(
			"permission <= ?", utils.GetFilePermissionsIndex(permissions),
		).Where(query, args...)

		return db.Where("id IN (?)", subQuery)
	}
}
