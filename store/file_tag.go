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

	if err := s.db.Scopes(scopes.Paginate(
		s.db.Scopes(
			filters.FilterByName(name),
			filters.FilterFileTagsByFileID(fileID),
		).Model(models.Tag{}),
		after, before, first, last,
	)).Find(&fileTags).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file tags!")
	}

	pageInfo := utils.GetEmptyPageInfo()

	if len(fileTags) > 0 {
		for _, fileTag := range fileTags {
			fileTagConnection.Edges = append(fileTagConnection.Edges, &models.FileTagEdge{
				Cursor: utils.EncodeCursor(fileTag.ID),
				Node:   fileTag,
			})
		}

		pageInfo.StartCursor = &fileTagConnection.Edges[0].Cursor
		pageInfo.EndCursor = &fileTagConnection.Edges[len(fileTagConnection.Edges)-1].Cursor

		if err := s.db.Scopes(scopes.HasPreviousPage(
			s.db.Scopes(
				filters.FilterByName(name),
				filters.FilterFileTagsByFileID(fileID),
			).Model(models.Tag{}),
			fileTags[0].ID,
		)).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(scopes.HasNextPage(
			s.db.Scopes(
				filters.FilterByName(name),
				filters.FilterFileTagsByFileID(fileID),
			).Model(models.Tag{}),
			fileTags[len(fileTags)-1].ID,
		)).First(&models.Tag{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileTagConnection.PageInfo = &pageInfo

	return &fileTagConnection, nil
}

func (s *FileTagStore) DeleteFileTags(claims *auth.Claims, fileTagsInput models.FileTagsInput) ([]*models.FileTag, error) {
	if err := s.db.Scopes(filters.FilterFileCollaboratorsByFilePermissions(
		models.FilePermissionsMaintain,
		"file_id = ? AND (collaborator_id = ? OR ? = ?)",
		fileTagsInput.FileID, claims.ID, claims.Kind, models.UserKindAdmin,
	)).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	var fileTags []*models.FileTag

	if err := s.db.Where("file_id = ? AND tag_name IN ?", fileTagsInput.FileID, fileTagsInput.TagNames).Find(&fileTags).Delete(&fileTags).Error; err != nil {
		return nil, gqlerror.Errorf("File tags not found!")
	}

	return fileTags, nil
}
