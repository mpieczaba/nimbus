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

func (s *Store) GetFile(query interface{}, args ...interface{}) (*File, error) {
	var file File

	if err := s.db.Where(query, args...).First(&file).Error; err != nil {
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

func (s *Store) GetAllFilesWithCondition(query interface{}, args ...interface{}) ([]*File, error) {
	var files []*File

	if err := s.db.Where(query, args...).Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}

func (s *Store) CreateFile(file *File) (*File, error) {
	if err := s.db.Create(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	return file, nil
}

func (s *Store) UpdateFile(file *File) (*File, error) {
	if err := s.db.Save(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	return file, nil
}

func (s *Store) DeleteFile(query interface{}, args ...interface{}) (*File, error) {
	var file File

	if err := s.db.Where(query, args...).First(&file).Select("Tags").Delete(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}
