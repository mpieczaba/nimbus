package tag

import (
	"github.com/mpieczaba/nimbus/database"
	"github.com/mpieczaba/nimbus/file"
)

type Tag struct {
	database.Model
	ID      string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Name    string `json:"name" gorm:"type:varchar(64);unique;not null"`
	OwnerID string `json:"ownerId" gorm:"type:varchar(20);not null"`
	Files   []file.FileTag
}

type TagInput struct {
	Name string `json:"name" validate:"required,tagname,min=3,max=64"`
}

type TagUpdateInput struct {
	Name    string `json:"name" validate:"omitempty,tagname,min=3,max=64"`
	OwnerID string `json:"ownerId" validate:"omitempty,alphanum,len=20"`
}
