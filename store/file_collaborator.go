package store

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/paginator"

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

func (s *FileCollaboratorStore) GetFileCollaborators(after, before *string, first, last *int, fileId string) (*models.FileCollaboratorConnection, error) {
	var fileCollaboratorConnection models.FileCollaboratorConnection
	var fileCollaborators []*models.User

	var permissions []models.FilePermission

	subQuery := s.db.Model(models.FileCollaborator{}).Where("file_id = ?", fileId).Select("permission").Find(&permissions).Select("collaborator_id")

	if err := s.db.Model(models.User{}).Where("id IN (?)", subQuery).Scopes(paginator.Paginate(after, before, first, last)).Find(&fileCollaborators).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file collaborators!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(fileCollaborators) > 0 {
		fileCollaboratorConnection.Nodes = fileCollaborators

		for i, fileCollaborator := range fileCollaborators {
			cursor, err := paginator.EncodeCursor(fileCollaborator.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all file collaborators!")
			}

			fileCollaboratorConnection.Edges = append(fileCollaboratorConnection.Edges, &models.FileCollaboratorEdge{
				Cursor:     cursor,
				Node:       fileCollaborator,
				Permission: permissions[i],
			})
		}

		if err := s.db.Model(models.User{}).Where("id IN (?)", subQuery).Scopes(paginator.GetBefore(fileCollaborators[0].ID)).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Model(models.User{}).Where("id IN (?)", subQuery).Scopes(paginator.GetAfter(fileCollaborators[len(fileCollaborators)-1].ID)).First(&models.User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileCollaboratorConnection.PageInfo = &pageInfo

	return &fileCollaboratorConnection, nil
}

func (s *FileCollaboratorStore) AddFileCollaborator(fileCollaborator *models.FileCollaborator) (*models.FileCollaborator, error) {
	if err := s.db.Create(fileCollaborator).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file collaborator already exists!")
	}

	return fileCollaborator, nil
}
