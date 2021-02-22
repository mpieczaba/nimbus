package models

import (
	"github.com/mpieczaba/nimbus/database"

	"github.com/99designs/gqlgen/graphql"
)

type File struct {
	database.Model
	ID        string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	MimeType  string `json:"mimeType" gorm:"type:varchar(127);not null"`
	Extension string `json:"extension" gorm:"type:varchar(10);not null"`
	Size      int64  `json:"size" gorm:"type:bigint"`
	OwnerID   string `json:"ownerId" gorm:"type:varchar(20);not null"`
}

type FileInput struct {
	Name      string           `json:"name"  validate:"required,filename,min=1,max=255"`
	Tags      []string         `json:"tags" validate:"required,dive,alphanum,len=20"`
	SharedFor []FileShareInput `validate:"omitempty,dive"`
	File      graphql.Upload   `json:"file"  validate:"required"`
}

type FileUpdateInput struct {
	Name      string           `json:"name" validate:"omitempty,filename,min=1,max=255"`
	OwnerID   string           `json:"ownerId" validate:"omitempty,alphanum,len=20"`
	Tags      []string         `json:"tags" validate:"omitempty,dive,alphanum,len=20"`
	SharedFor []FileShareInput `validate:"omitempty,dive"`
	File      graphql.Upload   `json:"file" validate:"omitempty"`
}

type FileTag struct {
	FileID string `json:"fileId" gorm:"type:varchar(20);not null"`
	TagID  string `json:"tagId" gorm:"type:varchar(20);not null"`
}

type FileShare struct {
	FileID      string `json:"fileId" gorm:"type:varchar(20);not null"`
	UserID      string `json:"userId" gorm:"type:varchar(20);not null"`
	Permissions int    `json:"permissions" gorm:"type:int;not null"`
}

type FileShareInput struct {
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}
