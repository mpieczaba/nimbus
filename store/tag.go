package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/utils"

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

func (s *TagStore) GetTag(query interface{}, args ...interface{}) (*models.Tag, error) {
	var tag models.Tag

	if err := s.db.Where(query, args...).First(&tag).Error; err != nil {
		return nil, gqlerror.Errorf("Tag not found!")
	}

	return &tag, nil
}

func (s *TagStore) GetAllTags(after, before *string, first, last *int, name *string) (*models.TagConnection, error) {
	var tagConnection models.TagConnection
	var tags []*models.Tag

	if err := s.db.Scopes(
		scopes.NameLike(models.Tag{}, "name", name),
		scopes.Paginate(after, before, first, last),
	).Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all tags!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(tags) > 0 {
		tagConnection.Nodes = tags

		for _, tag := range tags {
			cursor, err := utils.EncodeCursor(tag.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all tags!")
			}

			tagConnection.Edges = append(tagConnection.Edges, &models.TagEdge{
				Cursor: cursor,
				Node:   tag,
			})
		}

		if err := s.db.Scopes(
			scopes.NameLike(models.Tag{}, "name", name),
			scopes.GetBefore(tags[0].ID),
		).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(
			scopes.NameLike(models.Tag{}, "name", name),
			scopes.GetAfter(tags[len(tags)-1].ID),
		).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	tagConnection.PageInfo = &pageInfo

	return &tagConnection, nil
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
