package models

type FileTag struct {
	FileID string `json:"fileId" gorm:"type:varchar(20);not null"`
	TagID  string `json:"tagId" gorm:"type:varchar(20);not null"`
}
