package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type TagStore struct {
	db *gorm.DB
}

func NewTagStore(db *gorm.DB) *TagStore {
	return &TagStore{
		db: db,
	}
}

func (s *TagStore) CreateTagsOrAppendFileTags(claims *auth.Claims, tags []*models.Tag) ([]*models.Tag, error) {
	if err := s.db.Scopes(
		scopes.FilePermission(models.User{}, "collaborator_id", models.FilePermissionMaintain, "file_id = ? AND (collaborator_id = ? OR ? = ?)", tags[0].FileTags[0].FileID, claims.ID, claims.Kind, models.UserKindAdmin),
	).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	if err := s.db.Save(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file tag already exists!")
	}

	return tags, nil
}
