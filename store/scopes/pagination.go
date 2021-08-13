package scopes

import (
	"fmt"

	"github.com/mpieczaba/nimbus/utils"

	"gorm.io/gorm"
)

func Paginate(query *gorm.DB, after, before *string, first, last *int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc").Table("(?) as pag", db.Scopes(
			getLimit(first, last), getAfter(after), getBefore(before),
		).Table("(?) as tab", query))
	}
}

func getLimit(first, last *int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// "Strongly discouraged" - https://relay.dev/graphql/connections.htm#sel-FAJLFGBEBY22c
		if first != nil && last != nil {
			db.AddError(fmt.Errorf("'first' and 'last' cannot be given at the same time!"))

			return db
		}

		if first != nil && *first >= 0 && *first <= 50 {
			return db.Order("id desc").Limit(*first)
		} else if last != nil && *last >= 0 && *last <= 50 {
			return db.Order("id asc").Limit(*last)
		}

		return db
	}
}

func getAfter(after *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if after == nil {
			return db
		}

		decodedAfter, err := utils.DecodeCursor(*after)

		if err != nil {
			db.AddError(fmt.Errorf("Invalid 'after' cursor!"))

			return db
		}

		return db.Where("id < ?", decodedAfter)
	}
}

func getBefore(before *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if before == nil {
			return db
		}

		decodedBefore, err := utils.DecodeCursor(*before)

		if err != nil {
			db.AddError(fmt.Errorf("Invalid 'before' cursor!"))

			return db
		}

		return db.Where("id > ?", decodedBefore)
	}
}

func HasPreviousPage(query *gorm.DB, id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc").Where("id > ?", id).Limit(1).Table("(?) as p", query)
	}
}

func HasNextPage(query *gorm.DB, id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc").Where("id < ?", id).Limit(1).Table("(?) as p", query)
	}
}
