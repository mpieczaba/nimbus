package store

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type FileTagStore struct {
	db *gorm.DB
}

func NewFileTagStore(db *gorm.DB) *FileTagStore {
	return &FileTagStore{
		db: db,
	}
}

func (s *FileTagStore) GetFileTags(after, before *string, first, last *int, fileID string, name *string) (*models.FileTagConnection, error) {
	var fileTagConnection models.FileTagConnection
	var fileTags []*models.Tag

	if err := s.db.Scopes(
		scopes.FileTag(models.Tag{}, "tag_name", "name", "file_id = ?", fileID),
		scopes.NameLike(models.Tag{}, "name", name),
		scopes.Paginate(after, before, first, last),
	).Find(&fileTags).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file tags!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(fileTags) > 0 {
		fileTagConnection.Nodes = fileTags

		for _, fileTag := range fileTags {
			cursor, err := utils.EncodeCursor(fileTag.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all file tags!")
			}

			fileTagConnection.Edges = append(fileTagConnection.Edges, &models.FileTagEdge{
				Cursor: cursor,
				Node:   fileTag,
			})
		}

		if err := s.db.Scopes(
			scopes.FileTag(models.Tag{}, "tag_name", "name", "file_id = ?", fileID),
			scopes.NameLike(models.Tag{}, "name", name),
			scopes.GetBefore(fileTags[0].ID),
		).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(
			scopes.FileTag(models.Tag{}, "tag_name", "name", "file_id = ?", fileID),
			scopes.NameLike(models.Tag{}, "name", name),
			scopes.GetAfter(fileTags[len(fileTags)-1].ID),
		).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileTagConnection.PageInfo = &pageInfo

	return &fileTagConnection, nil
}
