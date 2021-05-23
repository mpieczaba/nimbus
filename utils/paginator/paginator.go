package paginator

import (
	"fmt"

	"gorm.io/gorm"
)

type queryConfig struct {
	limit  int
	after  string
	before string
	order  string
}

func Paginate(after, before *string, first, last *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		cfg, err := getQueryConfig(after, before, first, last)

		if err != nil {
			db.AddError(err)

			return db
		}

		subQuery := db.Session(&gorm.Session{PrepareStmt: true}).Order("id " + cfg.order).Limit(cfg.limit)

		if cfg.after != "" {
			subQuery.Where("id < ?", cfg.after)
		}

		if cfg.before != "" {
			subQuery.Where("id > ?", cfg.before)
		}

		return db.Order("id desc").Table("(?) as u", subQuery)
	}
}

func GetBefore(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fmt.Println(id)

		return db.Order("id desc").Where("id > ?", id).Limit(1)
	}
}

func GetAfter(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc").Where("id < ?", id).Limit(1)
	}
}

func getQueryConfig(after, before *string, first, last *int) (*queryConfig, error) {
	var cfg queryConfig

	// Set default limit
	cfg.limit = 50

	if after != nil && *after != "" {
		decodedAfter, err := DecodeCursor(*after)

		if err != nil {
			return nil, fmt.Errorf("Invalid 'after' cursor!")
		}

		cfg.after = decodedAfter
	}

	if before != nil && *before != "" {
		decodedBefore, err := DecodeCursor(*before)

		if err != nil {
			return nil, fmt.Errorf("Invalid 'before' cursor!")
		}

		cfg.before = decodedBefore
	}

	// "Strongly discouraged" - https://relay.dev/graphql/connections.htm#sel-FAJLFGBEBY22c
	if first != nil && last != nil {
		return nil, fmt.Errorf("'first' and 'last' cannot be given at the same time!")
	}

	if first != nil && *first >= 0 {
		cfg.limit = *first
		cfg.order = "desc"
	} else if last != nil && *last >= 0 {
		cfg.limit = *last
		cfg.order = "asc"
	}

	return &cfg, nil
}
