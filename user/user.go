package user

import (
	"time"

	"github.com/mpieczaba/nimbus/api/models"
)

type User struct {
	ID        string          `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Username  string          `json:"username" gorm:"type:varchar(64);unique;not null"`
	Password  string          `gorm:"type:varchar(128);not null"`
	Kind      models.UserKind `json:"kind" gorm:"type:varchar(8);not null"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *time.Time      `json:"deletedAt"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Kind     models.UserKind `json:"kind"`
}
