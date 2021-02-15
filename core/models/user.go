package models

import "github.com/mpieczaba/nimbus/core/database"

type User struct {
	database.Model
	ID       string `json:"id"  gorm:"type:varchar(20);primary_key;not null"`
	Username string `json:"username" gorm:"type:varchar(64);UNIQUE_INDEX;not null"`
	Password string `gorm:"type:varchar(128);not null"`
}
