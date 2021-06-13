package scopes

import (
	"gorm.io/gorm"
)

func NameLike(model interface{}, fieldToSelect string, stringToCompare *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stringToCompare == nil {
			return db.Model(model)
		}

		return db.Model(model).Where(fieldToSelect+" LIKE ?", *stringToCompare)
	}
}
