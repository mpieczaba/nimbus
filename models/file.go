package models

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type File struct {
	ID            string `json:"id" gorm:"type:varchar(20);primaryKey;not null"`
	Name          string `json:"name" gorm:"type:varchar(255);not null"`
	MimeType      string `json:"mimeType" gorm:"type:varchar(127);not null"`
	Extension     string `json:"extension" gorm:"type:varchar(10);not null"`
	Size          int64  `json:"size" gorm:"type:bigint; not null"`
	FileTags      []FileTag
	Collaborators []FileCollaborator
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

type FileInput struct {
	Name string         `json:"name" validate:"omitempty,filename,min=1,max=255"`
	File graphql.Upload `json:"file" validate:"required"`
}

type FileUpdateInput struct {
	Name string         `json:"name" validate:"omitempty,filename,min=1,max=255"`
	File graphql.Upload `json:"file" validate:"omitempty"`
}
