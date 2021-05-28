package models

import (
	"time"
)

type FileCollaborator struct {
	FileID         string         `json:"fileId" gorm:"foreignKey;uniqueIndex:file_collaborator;not null"`
	CollaboratorID string         `json:"collaboratorId" gorm:"foreignKey;uniqueIndex:file_collaborator;not null"`
	Permission     FilePermission `json:"permission" gorm:"type:varchar(8);not null"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      *time.Time     `json:"deletedAt"`
}

type FileCollaboratorInput struct {
	FileID         string         `json:"fileId"`
	CollaboratorID string         `json:"collaboratorId"`
	Permission     FilePermission `json:"permission"`
}
