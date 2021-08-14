package filters

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/utils"

	"gorm.io/gorm"
)

func FilterFilesByFilePermissions(claims *auth.Claims, permissions models.FilePermissions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subQuery := db.Session(&gorm.Session{NewDB: true}).Model(models.FileCollaborator{}).Select("file_id").Where(
			"permissions <= ?", utils.GetFilePermissionsIndex(permissions),
		).Where("collaborator_id = ? OR ? = ?", claims.ID, claims.Kind, models.UserKindAdmin)

		return db.Where("id IN (?)", subQuery)
	}
}

func FilterFilesByTags(tags []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(tags) == 0 {
			return db
		}

		subQuery := db.Session(&gorm.Session{NewDB: true}).Model(models.FileTag{}).Select("file_id").Where("tag_name IN ?", tags)

		return db.Where("id IN (?)", subQuery)
	}
}
