package tag_share

import "github.com/mpieczaba/nimbus/database"

type TagShare struct {
	database.Model
	ID          string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	TagID       string `json:"tagId" gorm:"type:varchar(20);not null"`
	UserID      string `json:"userId" gorm:"type:varchar(20);not null"`
	Permissions int    `json:"permissions" gorm:"type:int;not null"`
}

type TagShareInput struct {
	TagID       string `json:"tagId" validate:"required,alphanum,len=20"`
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}

type TagShareUpdateInput struct {
	Permissions int `json:"permissions" validate:"omitempty"`
}
