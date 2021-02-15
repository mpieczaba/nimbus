package models

import "github.com/mpieczaba/nimbus/core/database"

type User struct {
	database.Model
	ID       string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Username string `json:"username" gorm:"type:varchar(64);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
}

type UserInput struct {
	Username string `json:"username" validate:"required,min=3,max=64,lowercase"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateInput struct {
	Username string `json:"username" validate:"omitempty,min=3,max=64,lowercase"`
	Password string `json:"password" validate:"omitempty,min=8"`
}
