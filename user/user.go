package user

import "time"

type User struct {
	ID        string     `json:"id"  gorm:"type:varchar(20);primaryKey;not null"`
	Username  string     `json:"username" gorm:"type:varchar(64);unique;not null"`
	Password  string     `gorm:"type:varchar(128);not null"`
	Kind      UserKind   `json:"kind" gorm:"type:varchar(8);not null"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
