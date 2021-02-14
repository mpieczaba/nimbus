package models

import "github.com/mpieczaba/nimbus/core/database"

type File struct {
	database.Model
	ID   string `json:"id"  gorm:"type:varchar(20);primary_key;not null"`
	Name string `json:"name" gorm:"type:varchar(128);not null"`
}

type FileInput struct {
	Name string `json:"name"`
}

type FileUpdateInput struct {
	Name string `json:"name"`
}
