package file_tag

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

func (s *Store) GetFileTag(query interface{}, args ...interface{}) (*FileTag, error) {
	var fileTag FileTag

	if err := s.db.Where(query, args...).First(&fileTag).Error; err != nil {
		return nil, gqlerror.Errorf("File tag not found!")
	}

	return &fileTag, nil
}

func (s *Store) DeleteFileTag(query interface{}, args ...interface{}) (*FileTag, error) {
	var fileTag FileTag

	if err := s.db.Where(query, args...).Find(&fileTag).Delete(&fileTag).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete file tag!")
	}

	return &fileTag, nil
}

func (s *Store) GetTagIDs(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Select("tag_id").Where(query, args...).Table("file_tags")
}
