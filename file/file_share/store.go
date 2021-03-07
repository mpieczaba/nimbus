package file_share

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

func (s *Store) GetFileShare(query interface{}, args ...interface{}) (*FileShare, error) {
	var fileShare FileShare

	if err := s.db.Where(query, args...).First(&fileShare).Error; err != nil {
		return nil, gqlerror.Errorf("File share not found!")
	}

	return &fileShare, nil
}

func (s *Store) GetAllFileShares(query interface{}, args ...interface{}) ([]*FileShare, error) {
	var fileShares []*FileShare

	if err := s.db.Where(query, args...).Find(&fileShares).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting file shares!")
	}

	return fileShares, nil
}

func (s *Store) CreateFileShare(fileShare *FileShare) (*FileShare, error) {
	if err := s.db.Create(fileShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot save file share!")
	}

	return fileShare, nil
}

func (s *Store) UpdateFileShare(fileShare *FileShare) (*FileShare, error) {
	if err := s.db.Save(fileShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot update file share!")
	}

	return fileShare, nil
}

func (s *Store) DeleteFileShare(id string) (*FileShare, error) {
	var fileShare FileShare

	if err := s.db.Where("id = ?", id).Find(&fileShare).Delete(&fileShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete file share!")
	}

	return &fileShare, nil
}

func (s *Store) DeleteFileShares(query interface{}, args ...interface{}) ([]*FileShare, error) {
	var fileShares []*FileShare

	if err := s.db.Where(query, args...).Find(&fileShares).Delete(&fileShares).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete file shares!")
	}

	return fileShares, nil
}
