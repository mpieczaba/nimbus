package file_share

import "github.com/mpieczaba/nimbus/database"

type FileShare struct {
	database.Model
	FileID      string `json:"fileId" gorm:"foreignKey;uniqueIndex:file_share;not null"`
	UserID      string `json:"userId" gorm:"foreignKey;uniqueIndex:file_share;not null"`
	Permissions int    `json:"permissions" gorm:"type:tinyint;not null"`
}

type FileShareInput struct {
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}
