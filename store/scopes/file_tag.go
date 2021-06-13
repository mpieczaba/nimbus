package scopes

import (
	"github.com/mpieczaba/nimbus/models"

	"gorm.io/gorm"
)

func FileTag(model interface{}, fieldToSelect, keyField string, query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subQuery := db.Model(models.FileTag{}).Select(fieldToSelect).Where(query, args...)

		return db.Model(model).Where(keyField+" IN (?)", subQuery)
	}
}
