package models

import "github.com/mpieczaba/nimbus/core/database"

type Tag struct {
	database.Model
	ID      string `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Name    string `json:"name" gorm:"type:varchar(64);unique;not null"`
	OwnerID string `json:"ownerId" gorm:"type:varchar(20);not null"`
}
