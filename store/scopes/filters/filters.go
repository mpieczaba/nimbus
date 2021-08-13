package filters

import "gorm.io/gorm"

func FilterByName(name *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == nil {
			return db
		}

		return db.Where("name LIKE ?", *name)
	}
}

func FilterByUsername(username *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if username == nil {
			return db
		}

		return db.Where("username LIKE ?", *username)
	}
}
