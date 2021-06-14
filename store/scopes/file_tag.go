package scopes

import (
	"github.com/mpieczaba/nimbus/models"

	"gorm.io/gorm"
)

func FileTag(model interface{}, fieldToSelect, keyField string, query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s, ok := args[0].([]string); len(args) == 0 || (ok && len(s) == 0) {
			return db.Model(model)
		}

		subQuery := db.Model(models.FileTag{}).Select(fieldToSelect).Where(query, args...)

		return db.Session(&gorm.Session{NewDB: true}).Model(model).Where(keyField+" IN (?)", subQuery)
	}
}
