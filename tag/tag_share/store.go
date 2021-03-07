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

func (s *Store) CreateTagShare(tagShare *TagShare) (*TagShare, error) {
	if err := s.db.Save(tagShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot create tag share!")
	}

	return tagShare, nil
}

func (s *Store) UpdateTagShare(tagShare *TagShare) (*TagShare, error) {
	if err := s.db.Save(tagShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot update tag share!")
	}

	return tagShare, nil
}

func (s *Store) DeleteTagShare(id string) (*TagShare, error) {
	var tagShare TagShare

	if err := s.db.Where("id = ?", id).Find(&tagShare).Delete(&tagShare).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete tag share!")
	}

	return &tagShare, nil
}

func (s *Store) DeleteTagShares(query interface{}, args ...interface{}) ([]*TagShare, error) {
	var tagShares []*TagShare

	if err := s.db.Where(query, args...).Find(&tagShares).Delete(&tagShares).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete tag share!")
	}

	return tagShares, nil
}
