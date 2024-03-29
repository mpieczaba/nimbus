package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type FileCollaboratorStore struct {
	db *gorm.DB
}

func NewFileCollaboratorStore(db *gorm.DB) *FileCollaboratorStore {
	return &FileCollaboratorStore{
		db: db,
	}
}

func (s *FileCollaboratorStore) GetFileCollaborators(after, before *string, first, last *int, fileID string, username *string, permission models.FilePermission) (*models.FileCollaboratorConnection, error) {
	var fileCollaboratorConnection models.FileCollaboratorConnection
	var fileCollaborators []*models.User

	if err := s.db.Scopes(
		scopes.FilePermission(models.User{}, "collaborator_id", permission, "file_id = ?", fileID),
		scopes.NameLike(models.User{}, "username", username),
		scopes.Paginate(after, before, first, last),
	).Find(&fileCollaborators).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file collaborators!")
	}

	var permissions []int8

	if err := s.db.Model(models.FileCollaborator{}).Where("file_id = ? AND permission <= ?", fileID, utils.GetFilePermissionIndex(permission)).Select("permission").Find(&permissions).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file collaborators!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(fileCollaborators) > 0 {
		fileCollaboratorConnection.Nodes = fileCollaborators

		for i, fileCollaborator := range fileCollaborators {
			cursor, err := utils.EncodeCursor(fileCollaborator.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all file collaborators!")
			}

			fileCollaboratorConnection.Edges = append(fileCollaboratorConnection.Edges, &models.FileCollaboratorEdge{
				Cursor:     cursor,
				Node:       fileCollaborator,
				Permission: models.AllFilePermission[permissions[i]],
			})
		}

		if err := s.db.Scopes(
			scopes.FilePermission(models.User{}, "collaborator_id", permission, "file_id = ?", fileID),
			scopes.NameLike(models.User{}, "username", username),
			scopes.GetBefore(fileCollaborators[0].ID),
		).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(
			scopes.FilePermission(models.User{}, "collaborator_id", permission, "file_id = ?", fileID),
			scopes.NameLike(models.User{}, "username", username),
			scopes.GetAfter(fileCollaborators[len(fileCollaborators)-1].ID),
		).First(&models.User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileCollaboratorConnection.PageInfo = &pageInfo

	return &fileCollaboratorConnection, nil
}

func (s *FileCollaboratorStore) CreateFileCollaborator(claims *auth.Claims, fileCollaborator *models.FileCollaborator) (*models.FileCollaborator, error) {
	if err := s.db.Scopes(
		scopes.FilePermission(models.User{}, "collaborator_id", models.FilePermissionAdmin, "file_id = ? AND (collaborator_id = ? OR ? = ?)", fileCollaborator.FileID, claims.ID, claims.Kind, models.UserKindAdmin),
	).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	if err := s.db.Create(fileCollaborator).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file collaborator already exists!")
	}

	return fileCollaborator, nil
}

func (s *FileCollaboratorStore) DeleteFileCollaborator(claims *auth.Claims, fileID, collaboratorID string) (*models.FileCollaborator, error) {
	if err := s.db.Scopes(
		scopes.FilePermission(models.User{}, "collaborator_id", models.FilePermissionAdmin, "file_id = ? AND (collaborator_id = ? OR ? = ?)", fileID, claims.ID, claims.Kind, models.UserKindAdmin),
	).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	var fileCollaborator models.FileCollaborator

	if err := s.db.Where("file_id = ? AND collaborator_id = ?", fileID, collaboratorID).First(&fileCollaborator).Delete(&fileCollaborator).Error; err != nil {
		return nil, gqlerror.Errorf("File collaborator not found!")
	}

	return &fileCollaborator, nil
}
