package models

import "github.com/mpieczaba/nimbus/database"

type Tag struct {
	database.Model
	ID      string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Name    string `json:"name" gorm:"type:varchar(64);unique;not null"`
	OwnerID string `json:"ownerId" gorm:"type:varchar(20);not null"`
}

type TagInput struct {
	Name      string          `json:"name" validate:"required,tagname,min=3,max=64"`
	SharedFor []TagShareInput `validate:"omitempty,dive"`
}

type TagUpdateInput struct {
	Name      string          `json:"name" validate:"omitempty,tagname,min=3,max=64"`
	OwnerID   string          `json:"ownerId" validate:"omitempty,alphanum,len=20"`
	SharedFor []TagShareInput `validate:"omitempty,dive"`
}

type TagShare struct {
	TagID       string `json:"tagId" gorm:"type:varchar(20);not null"`
	UserID      string `json:"userId" gorm:"type:varchar(20);not null"`
	Permissions int    `json:"permissions" gorm:"type:int;not null"`
}

type TagShareInput struct {
	UserID      string `json:"userId" validate:"required,alphanum,len=20"`
	Permissions int    `json:"permissions" validate:"required"`
}
