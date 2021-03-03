package file_share

import (
	"github.com/mpieczaba/nimbus/database"
)

type FileShare struct {
	database.Model
	ID          string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	FileID      string `json:"fileId" gorm:"type:varchar(20);not null"`
	UserID      string `json:"userId" gorm:"type:varchar(20);not null"`
	Permissions int    `json:"permissions" gorm:"type:int;not null"`
}

type FileShareInput struct {
	FileID      string `json:"fileId" validate:"required,alphanum,len=20"`
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}

type FileShareUpdateInput struct {
	Permissions int `json:"permissions" validate:"omitempty"`
}
