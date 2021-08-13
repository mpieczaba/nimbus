package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/store/scopes/filters"
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

	if err := s.db.Scopes(scopes.Paginate(
		s.db.Scopes(filters.FilterByName(name)).Model(models.Tag{}),
		after, before, first, last,
	)).Find(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all tags!")
	}

	pageInfo := utils.GetEmptyPageInfo()

	if len(tags) > 0 {
		for _, tag := range tags {
			tagConnection.Edges = append(tagConnection.Edges, &models.TagEdge{
				Cursor: utils.EncodeCursor(tag.ID),
				Node:   tag,
			})
		}

		pageInfo.StartCursor = &tagConnection.Edges[0].Cursor
		pageInfo.EndCursor = &tagConnection.Edges[len(tagConnection.Edges)-1].Cursor

		if err := s.db.Scopes(scopes.HasPreviousPage(
			s.db.Scopes(filters.FilterByName(name)).Model(models.Tag{}),
			tags[0].ID,
		)).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(scopes.HasNextPage(
			s.db.Scopes(filters.FilterByName(name)).Model(models.Tag{}),
			tags[len(tags)-1].ID,
		)).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	tagConnection.PageInfo = &pageInfo

	return &tagConnection, nil
}

func (s *TagStore) CreateTagsOrAppendFileTags(claims *auth.Claims, tags []*models.Tag) ([]*models.Tag, error) {
	if err := s.db.Scopes(filters.FilterFileCollaboratorsByFilePermissions(
		models.FilePermissionsMaintain,
		"file_id = ? AND (collaborator_id = ? OR ? = ?)",
		tags[0].FileTags[0].FileID, claims.ID, claims.Kind, models.UserKindAdmin,
	)).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	if err := s.db.Save(&tags).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file tag already exists!")
	}

	return tags, nil
}
