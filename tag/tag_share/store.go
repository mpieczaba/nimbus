package tag_share

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

func (s *Store) GetTagShare(query interface{}, args ...interface{}) (*TagShare, error) {
	var tagShare TagShare

	if err := s.db.Where(query, args...).First(&tagShare).Error; err != nil {
		return nil, gqlerror.Errorf("Tag share not found!")
	}

	return &tagShare, nil
}

func (s *Store) GetAllTagShares(query interface{}, args ...interface{}) ([]*TagShare, error) {
	var tagShares []*TagShare

	if err := s.db.Where(query, args...).Find(&tagShares).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting tag shares!")
	}

	return tagShares, nil
}

func (s *Store) DeleteTagShare(query interface{}, args ...interface{}) (*TagShare, error) {
	var tagShare TagShare

	if err := s.db.Where(query, args...).Find(&tagShare).Delete(&tagShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete tag share!")
	}

	return &tagShare, nil
}
