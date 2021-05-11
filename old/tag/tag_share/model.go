package tag_share

import "github.com/mpieczaba/nimbus/database"

type TagShare struct {
	database.Model
	TagID     string       `json:"tagId" gorm:"foreignKey;uniqueIndex:tag_share;not null"`
	UserID    string       `json:"userId" gorm:"foreignKey;uniqueIndex:tag_share;not null"`
	ShareKind TagShareKind `json:"shareKind" gorm:"type:varchar(8);not null"`
}

type TagShareInput struct {
	UserID    string       `json:"userId" validate:"required,alphanum,len=20"`
	ShareKind TagShareKind `json:"shareKind" validate:"required"`
}
