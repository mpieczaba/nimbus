package file

import (
	"time"

	"github.com/mpieczaba/nimbus/api/models"
)

type File struct {
	ID        string     `json:"id" gorm:"type:varchar(20);primaryKey;not null"`
	Name      string     `json:"name" gorm:"type:varchar(64);unique;not null"`
	MimeType  string     `json:"mimeType" gorm:"type:varchar(127);not null"`
	Extension string     `json:"extension" gorm:"type:varchar(10);not null"`
	Size      int64      `json:"size" gorm:"type:bigint; not null"`
	OwnerID   string     `json:"ownerId" gorm:"foreignKey;not null"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type FileConnection struct {
	Edges    []*FileEdge      `json:"edges"`
	Nodes    []*File          `json:"nodes"`
	PageInfo *models.PageInfo `json:"pageInfo"`
}

type FileEdge struct {
	Cursor string `json:"cursor"`
	Node   *File  `json:"node"`
}
