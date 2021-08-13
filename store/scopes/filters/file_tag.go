package filters

import (
	"github.com/mpieczaba/nimbus/models"
	"gorm.io/gorm"
)

func FilterFileTagsByFileID(fileID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subQuery := db.Session(&gorm.Session{NewDB: true}).Model(models.FileTag{}).Select("tag_name").Where("file_id = ?", fileID)

		return db.Where("id IN (?)", subQuery)
	}
}
