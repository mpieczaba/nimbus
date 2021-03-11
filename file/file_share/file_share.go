package file_share

import "github.com/mpieczaba/nimbus/database"

type FileShare struct {
	database.Model
	FileID    string        `json:"fileId" gorm:"foreignKey;uniqueIndex:file_share;not null"`
	UserID    string        `json:"userId" gorm:"foreignKey;uniqueIndex:file_share;not null"`
	ShareKind FileShareKind `json:"shareKind" gorm:"type:varchar(8);not null"`
}

type FileShareInput struct {
	UserID    string        `json:"userId" validate:"required,alphanum,len=20"`
	ShareKind FileShareKind `json:"shareKind" validate:"required"`
}
