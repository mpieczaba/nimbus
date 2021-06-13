package models

import "time"

type Tag struct {
	ID        string     `json:"id" gorm:"type:varchar(20);primaryKey;not null"`
	Name      string     `json:"name" gorm:"type:varchar(32);unique;not null"`
	FileTags  []FileTag  `gorm:"foreignKey:TagName;references:Name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
