package file

import (
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/file/file_tag"

	"github.com/99designs/gqlgen/graphql"
)

type File struct {
	database.Model
	ID         string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Name       string `json:"name" gorm:"type:varchar(255);not null"`
	MimeType   string `json:"mimeType" gorm:"type:varchar(127);not null"`
	Extension  string `json:"extension" gorm:"type:varchar(10);not null"`
	Size       int64  `json:"size" gorm:"type:bigint"`
	OwnerID    string `json:"ownerId" gorm:"foreignKey;not null"`
	FileTags   []file_tag.FileTag
	FileShares []file_share.FileShare
}

type FileInput struct {
	Name      string                      `json:"name"  validate:"required,filename,min=1,max=255"`
	Tags      []string                    `json:"tags" validate:"required,dive,alphanum,len=20"`
	SharedFor []file_share.FileShareInput `json:"sharedFor" validate:"omitempty,dive"`
	File      graphql.Upload              `json:"file"  validate:"required"`
}

type FileUpdateInput struct {
	Name      string                      `json:"name" validate:"omitempty,filename,min=1,max=255"`
	OwnerID   string                      `json:"ownerId" validate:"omitempty,alphanum,len=20"`
	Tags      []string                    `json:"tags" validate:"omitempty,dive,alphanum,len=20"`
	SharedFor []file_share.FileShareInput `json:"sharedFor" validate:"omitempty,dive"`
	File      graphql.Upload              `json:"file" validate:"omitempty"`
}
