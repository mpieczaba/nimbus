package file_share

type FileShare struct {
	FileID      string `json:"fileId" gorm:"foreignKey;not null"`
	UserID      string `json:"userId" gorm:"foreignKey;not null"`
	Permissions int    `json:"permissions" gorm:"type:int;not null"`
}

type FileShareInput struct {
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}
