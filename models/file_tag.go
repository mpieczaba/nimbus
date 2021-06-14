package models

type FileTag struct {
	FileID  string `json:"fileId" gorm:"foreignKey;uniqueIndex:file_tag;not null"`
	TagName string `json:"tagName" gorm:"foreignKey;uniqueIndex:file_tag;not null"`
}

type FileTagsInput struct {
	FileID   string   `json:"fileId" validate:"required,alphanum,len=20"`
	TagNames []string `json:"tagNames" validate:"required,dive,tagname,min=1,max=32"`
}
