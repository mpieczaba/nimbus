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

type UserConnection struct {
	Edges    []*UserEdge      `json:"edges"`
	Nodes    []*User          `json:"nodes"`
	PageInfo *models.PageInfo `json:"pageInfo"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type UserInput struct {
	Username string `json:"username" validate:"required,username,min=3,max=64,lowercase"`
	Password string `json:"password" validate:"required,password,min=8,max=512"`
}

type UserUpdateInput struct {
	Username string          `json:"username" validate:"omitempty,username,min=3,max=64,lowercase"`
	Password string          `json:"password" validate:"omitempty,password,min=8,max=512"`
	Kind     models.UserKind `json:"kind" validate:"omitempty"`
}
