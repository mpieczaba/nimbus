package tag

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

func (s *Store) GetTag(query interface{}, args ...interface{}) (*Tag, error) {
	var tag Tag

	if err := s.db.Where(query, args).First(&tag).Error; err != nil {
		return nil, gqlerror.Errorf("Tag not found!")
	}

	return &tag, nil
}

func (s *Store) GetAllTags() ([]*Tag, error) {
	var tags []*Tag

	if err := s.db.Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all tags!")
	}

	return tags, nil
}

func (s *Store) GetAllTagsWithCondition(query interface{}, args ...interface{}) ([]*Tag, error) {
	var tags []*Tag

	if err := s.db.Where(query, args).Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all tags!")
	}

	return tags, nil
}
