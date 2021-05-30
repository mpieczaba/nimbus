package models

import (
	"time"
)

type FileCollaborator struct {
	FileID         string     `json:"fileId" gorm:"foreignKey;uniqueIndex:file_collaborator;not null"`
	CollaboratorID string     `json:"collaboratorId" gorm:"foreignKey;uniqueIndex:file_collaborator;not null"`
	Permission     int8       `json:"permission" gorm:"type:tinyint;not null"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
}

type FileCollaboratorInput struct {
	FileID         string         `json:"fileId" validate:"required,alphanum,len=20"`
	CollaboratorID string         `json:"collaboratorId" validate:"required,alphanum,len=20"`
	Permission     FilePermission `json:"permission"`
}
