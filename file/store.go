package file

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	store := &Store{
		db: db,
	}

	return store
}

func (s *Store) GetFileById(id string) (*File, error) {
	var file File

	if err := s.db.Where("id = ?", id).First(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *Store) GetAllFiles() ([]*File, error) {
	var files []*File

	if err := s.db.Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}
