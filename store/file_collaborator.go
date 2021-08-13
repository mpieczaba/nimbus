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

type FileCollaboratorStore struct {
	db *gorm.DB
}

func NewFileCollaboratorStore(db *gorm.DB) *FileCollaboratorStore {
	return &FileCollaboratorStore{
		db: db,
	}
}

func (s *FileCollaboratorStore) GetFileCollaborators(after, before *string, first, last *int, fileID string, username *string, permission models.FilePermissions) (*models.FileCollaboratorConnection, error) {
	var fileCollaboratorConnection models.FileCollaboratorConnection

	type result struct {
		*models.User
		Permission int8
	}

	var results []result

	subQuery := s.db.Scopes(
		filters.FilterFileCollaboratorsByFilePermissions(permission, "file_id = ?", fileID),
		filters.FilterByUsername(username),
	).Model(models.User{})

	if err := s.db.Scopes(scopes.Paginate(
		s.db.Select("p.*", "permission").Table("(?) as p", subQuery).Joins(
			"LEFT JOIN file_collaborators fc on p.id = fc.collaborator_id AND fc.file_id = ?", fileID,
		).Model(result{}),
		after, before, first, last,
	)).Find(&results).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting file collaborators!")
	}

	pageInfo := utils.GetEmptyPageInfo()

	if len(results) > 0 {
		for _, fileCollaborator := range results {
			fileCollaboratorConnection.Edges = append(fileCollaboratorConnection.Edges, &models.FileCollaboratorEdge{
				Cursor:      utils.EncodeCursor(fileCollaborator.ID),
				Node:        fileCollaborator.User,
				Permissions: models.AllFilePermissions[fileCollaborator.Permission],
			})
		}

		pageInfo.StartCursor = &fileCollaboratorConnection.Edges[0].Cursor
		pageInfo.EndCursor = &fileCollaboratorConnection.Edges[len(fileCollaboratorConnection.Edges)-1].Cursor

		if err := s.db.Scopes(scopes.HasPreviousPage(
			s.db.Scopes(
				filters.FilterFileCollaboratorsByFilePermissions(permission, "file_id = ?", fileID),
				filters.FilterByUsername(username),
			).Model(models.User{}),
			results[0].ID,
		)).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(scopes.HasNextPage(
			s.db.Scopes(
				filters.FilterFileCollaboratorsByFilePermissions(permission, "file_id = ?", fileID),
				filters.FilterByUsername(username),
			).Model(models.User{}),
			results[len(results)-1].ID,
		)).First(&models.User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileCollaboratorConnection.PageInfo = &pageInfo

	return &fileCollaboratorConnection, nil
}

func (s *FileCollaboratorStore) CreateFileCollaborator(claims *auth.Claims, fileCollaborator *models.FileCollaborator) (*models.FileCollaborator, error) {
	if err := s.db.Scopes(filters.FilterFileCollaboratorsByFilePermissions(
		models.FilePermissionsAdmin,
		"file_id = ? AND (collaborator_id = ? OR ? = ?)",
		fileCollaborator.FileID, claims.ID, claims.Kind, models.UserKindAdmin,
	)).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	if err := s.db.Create(fileCollaborator).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file collaborator already exists!")
	}

	return fileCollaborator, nil
}

func (s *FileCollaboratorStore) DeleteFileCollaborator(claims *auth.Claims, fileID, collaboratorID string) (*models.FileCollaborator, error) {
	if err := s.db.Scopes(filters.FilterFileCollaboratorsByFilePermissions(
		models.FilePermissionsAdmin,
		"file_id = ? AND (collaborator_id = ? OR ? = ?)",
		fileID, claims.ID, claims.Kind, models.UserKindAdmin,
	)).First(&models.User{}).Error; err != nil {
		return nil, gqlerror.Errorf("No required permission!")
	}

	var fileCollaborator models.FileCollaborator

	if err := s.db.Where("file_id = ? AND collaborator_id = ?", fileID, collaboratorID).First(&fileCollaborator).Delete(&fileCollaborator).Error; err != nil {
		return nil, gqlerror.Errorf("File collaborator not found!")
	}

	return &fileCollaborator, nil
}
