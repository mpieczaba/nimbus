package tag_share

import "github.com/mpieczaba/nimbus/database"

type TagShare struct {
	database.Model
	TagID       string `json:"tagId" gorm:"foreignKey;uniqueIndex:tag_share;not null"`
	UserID      string `json:"userId" gorm:"foreignKey;uniqueIndex:tag_share;not null"`
	Permissions int    `json:"permissions" gorm:"type:tinyint;not null"`
}

type TagShareInput struct {
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}
