package file_tag

type FileTag struct {
	FileID string `json:"fileId" gorm:"foreignKey;uniqueIndex:file_tag;not null"`
	TagID  string `json:"tagId" gorm:"foreignKey;uniqueIndex:file_tag;not null"`
}
