package user

import (
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/tag/tag_share"
)

type User struct {
	database.Model
	ID         string      `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Username   string      `json:"username" gorm:"type:varchar(64);unique;not null"`
	Password   string      `gorm:"type:varchar(128);not null"`
	UserFiles  []file.File `gorm:"foreignKey:OwnerID"`
	FileShares []file_share.FileShare
	UserTags   []tag.Tag `gorm:"foreignKey:OwnerID"`
	TagShares  []tag_share.TagShare
}

type UserInput struct {
	Username string `json:"username" validate:"required,username,min=3,max=64,lowercase"`
	Password string `json:"password" validate:"required,password,min=8,max=512"`
}

type UserUpdateInput struct {
	Username string `json:"username" validate:"omitempty,username,min=3,max=64,lowercase"`
	Password string `json:"password" validate:"omitempty,password,min=8,max=512"`
}
