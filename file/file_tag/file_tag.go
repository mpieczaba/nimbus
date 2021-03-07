package file_tag

type FileTag struct {
	FileID string `json:"fileId" gorm:"foreignKey;not null"`
	TagID  string `json:"tagId" gorm:"foreignKey;not null"`
}
